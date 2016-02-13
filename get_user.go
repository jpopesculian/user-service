package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	user, err := RepoGetUserById(id)
	if err != nil {
		http.Error(w, "User Not Found!", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}
