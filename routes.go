package main

import (
	"net/http"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}

var routes = []Route{
	Route{
		"Authenticate",
		"GET",
		"/authenticated",
		Authenticate,
	},
	Route{
		"GetUser",
		"GET",
		"/{id}",
		GetUser,
	},
	Route{
		"Login",
		"POST",
		"/login",
		Login,
	},
	Route{
		"Register",
		"POST",
		"/register",
		Register,
	},
}
