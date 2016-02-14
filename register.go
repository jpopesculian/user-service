package main

import (
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var form RegisterForm
	if err := ReadJsonForm(r, &form); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user, err := RepoCreateUser(form.Email, form.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	accessToken, err := CreateUserToken(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	WriteJson(w, accessToken)
}
