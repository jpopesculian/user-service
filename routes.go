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
		"GetUser",
		"GET",
		"/user/{id}",
		GetUser,
	},
	Route{
		"Authenticate",
		"POST",
		"/authenticate",
		Authenticate,
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
