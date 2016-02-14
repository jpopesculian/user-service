package main

import (
	"regexp"
	"strings"
)

type EmailPassForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ValidateEmailPassForm(form EmailPassForm) (bool, error) {
	match, err := regexp.MatchString("^[^@]+@[^@]+\\.[^@]+$", form.Email)
	if err != nil || !match {
		return match, err
	}
	match, err = regexp.MatchString(".{6}.*", form.Password)
	if err != nil || !match {
		return match, err
	}
	return match, err
}

func TrimEmailPassForm(form EmailPassForm) EmailPassForm {
	form.Email = strings.TrimSpace(strings.ToLower(form.Email))
	form.Password = strings.TrimSpace(form.Password)
	return form
}
