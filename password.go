package main

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	var hashed string
	byteArray := []byte(password)
	output, err := bcrypt.GenerateFromPassword(byteArray, bcrypt.DefaultCost)
	if err != nil {
		return hashed, err
	}
	hashed = string(output[:])
	return hashed, nil
}

func ComparePassword(hashed string, password string) bool {
	hashedBytes := []byte(hashed)
	passBytes := []byte(password)
	err := bcrypt.CompareHashAndPassword(hashedBytes, passBytes)
	return err == nil
}
