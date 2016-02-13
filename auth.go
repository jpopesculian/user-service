package main

import (
	"fmt"
	"net/http"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
