package main

import (
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var form EmailPassForm
	if err := ReadJsonForm(r, &form); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	form = TrimEmailPassForm(form)
	id, ok, err := ValidLoginAttempt(form.Email, form.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "Bad Login Attempt!", http.StatusUnauthorized)
		return
	}
	user, err := RepoGetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	accessToken, err := CreateUserToken(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteAccessTokenToResponse(w, accessToken)
	WriteJson(w, user)
}
