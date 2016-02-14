package main

import (
	"net/http"
)

type AuthenticatedUser struct {
	User        User        `json:"user"`
	AccessToken AccessToken `json:"jwt"`
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	var accessToken AccessToken
	if err := ReadJsonForm(r, &accessToken); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	authenticatedUser := AuthenticatedUser{
		user,
		newAccessToken,
	}
	WriteJson(w, authenticatedUser)
}
