package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type AccessToken struct {
	Value string `json:"access_token"`
}

type AuthenticatedUser struct {
	User        User        `json:"user"`
	AccessToken AccessToken `json:"jwt"`
}

func CreateUserToken(id string) (AccessToken, error) {
	var accessToken AccessToken
	token := jwt.New(jwt.SigningMethodRS256)
	token.Claims["id"] = id
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	key, err := GetPrivateKey()
	if err != nil {
		return accessToken, err
	}
	tokenString, err := token.SignedString(key)
	if err != nil {
		return accessToken, err
	}
	accessToken = AccessToken{tokenString}
	return accessToken, nil
}

func GetUserIdFromToken(accessToken AccessToken) (string, error) {
	var id string
	token, err := jwt.Parse(accessToken.Value, GetSigningKeyFromToken)
	if err != nil {
		return id, err
	}
	if token.Valid {
		claim := token.Claims["id"]
		if id, ok := claim.(string); ok {
			return id, nil
		}
	}
	return id, fmt.Errorf("Invalid Token!")
}

func GetSigningKeyFromToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	key, err := GetPublicKey()
	if err != nil {
		return key, err
	}
	return key, nil
}
