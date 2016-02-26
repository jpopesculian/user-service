package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

const AccessTokenHeader = "X-Authorization"

func CreateUserToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	token.Claims["id"] = id
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	key, err := GetPrivateKey()
	if err != nil {
		return "", err
	}
	accessToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func GetUserIdFromToken(accessToken string) (string, error) {
	var id string
	token, err := jwt.Parse(accessToken, GetSigningKeyFromToken)
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

func GetAccessTokenFromRequest(r *http.Request) string {
	return r.Header.Get(AccessTokenHeader)
}

func WriteAccessTokenToResponse(w http.ResponseWriter, accessToken string) {
	w.Header().Set(AccessTokenHeader, accessToken)
}
