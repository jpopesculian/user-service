package main

import (
	"net/http"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	accessToken := GetAccessTokenFromRequest(r)
	id, err := GetUserIdFromToken(accessToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	user, err := RepoGetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	newAccessToken, err := CreateUserToken(user.Id)
	if err != nil {
		newAccessToken = accessToken
	}
	WriteAccessTokenToResponse(w, newAccessToken)
	WriteJson(w, user)
}
