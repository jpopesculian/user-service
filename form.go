package main

type EmailPassForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterForm EmailPassForm
type LoginForm EmailPassForm
