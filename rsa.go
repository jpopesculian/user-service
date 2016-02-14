package main

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
)

func GetPublicKey() (*rsa.PublicKey, error) {
	var key *rsa.PublicKey
	bytes, err := ioutil.ReadFile(config.PublicKeyPath)
	if err != nil {
		return key, err
	}
	return jwt.ParseRSAPublicKeyFromPEM(bytes)
}

func GetPrivateKey() (*rsa.PrivateKey, error) {
	var key *rsa.PrivateKey
	bytes, err := ioutil.ReadFile(config.PrivateKeyPath)
	if err != nil {
		return key, err
	}
	return jwt.ParseRSAPrivateKeyFromPEM(bytes)
}
