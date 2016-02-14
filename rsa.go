package main

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
)

func GetPublicKey() (*rsa.PublicKey, error) {
	var key *rsa.PublicKey
	privateKey, err := GetPrivateKey()
	if err != nil {
		return key, err
	}
	return &privateKey.PublicKey, nil
}

func GetPrivateKey() (*rsa.PrivateKey, error) {
	var key *rsa.PrivateKey
	bytes, err := ioutil.ReadFile(config.PrivateKeyPath)
	if err != nil {
		return key, err
	}
	return jwt.ParseRSAPrivateKeyFromPEM(bytes)
}
