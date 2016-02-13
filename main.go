package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	InitEtcd()
	router := NewRouter()
	config := NewConfig()
	log.Println(fmt.Sprintf("Serving at localhost:%d...", config.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
