package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	key := fmt.Sprintf("/users/data/%s/email", id)
	meta, err := etcd.Get(context.Background(), key, nil)
	if err != nil {
		http.Error(w, "User Not Found!", http.StatusNotFound)
		return
	}
	email := meta.Node.Value
	user := User{
		id,
		email,
	}
	json.NewEncoder(w).Encode(user)
}
