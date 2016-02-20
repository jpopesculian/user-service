package main

import (
	"errors"
	"fmt"
	"github.com/coreos/etcd/client"
	"github.com/nu7hatch/gouuid"
	"golang.org/x/net/context"
)

func RepoGetUserById(id string) (User, error) {
	var user User
	key := fmt.Sprintf("/users/data/%s/email", id)
	meta, err := etcd.Get(context.Background(), key, nil)
	if err != nil {
		return user, err
	}
	email := meta.Node.Value
	user = User{
		id,
		email,
	}
	return user, nil
}

func RepoCreateUser(email string, password string) (User, error) {
	var user User
	if ok, _, _ := RepoUserExistsByEmail(email); ok {
		return user, errors.New("Email Already Exists!")
	}
	id, err := RepoCreateUniqueId(3)
	if err != nil {
		return user, err
	}
	emailIndex := fmt.Sprintf("/users/indexes/email/%s", email)
	_, err = etcd.Create(context.Background(), emailIndex, id)
	if err != nil {
		return user, err
	}
	userEmailKey := fmt.Sprintf("/users/data/%s/email", id)
	_, err = etcd.Create(context.Background(), userEmailKey, email)
	if err != nil {
		return user, err
	}
	hashed, err := HashPassword(password)
	if err != nil {
		return user, err
	}
	userPasswordServiceKey := fmt.Sprintf("/users/services/%s/password", id)
	_, err = etcd.Create(context.Background(), userPasswordServiceKey, hashed)
	if err != nil {
		return user, err
	}
	user = User{
		id,
		email,
	}
	return user, nil
}

func RepoGetUserIdByEmail(email string) (string, error) {
	var id string
	key := fmt.Sprintf("/users/indexes/email/%s", email)
	meta, err := etcd.Get(context.Background(), key, nil)
	if err != nil {
		return id, err
	}
	id = meta.Node.Value
	return id, nil
}

func RepoUserExistsByEmail(email string) (bool, string, error) {
	userId, err := RepoGetUserIdByEmail(email)
	if err != nil {
		if client.IsKeyNotFound(err) != true {
			return false, userId, err
		}
		return false, userId, nil
	}
	return true, userId, nil
}

func RepoGetUserPasswordById(id string) (string, error) {
	var password string
	key := fmt.Sprintf("/users/services/%s/password", id)
	meta, err := etcd.Get(context.Background(), key, nil)
	if err != nil {
		return id, err
	}
	password = meta.Node.Value
	return password, nil
}

func RepoCreateUniqueId(tries int) (string, error) {
	var id string
	try := 0
	for try < tries {
		try++
		uuid, err := uuid.NewV4()
		if err != nil {
			continue
		}
		id = uuid.String()
		_, err = RepoGetUserById(id)
		if err != nil && client.IsKeyNotFound(err) {
			return id, nil
		}
	}
	return id, errors.New("Unable to Create UUID!")
}
