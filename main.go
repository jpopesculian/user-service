package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	init_etcd()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":1337", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
