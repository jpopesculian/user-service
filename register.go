package main

import (
	"encoding/json"
	"net/http"
)

type RegisterForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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
	json.NewEncoder(w).Encode(user)
}
