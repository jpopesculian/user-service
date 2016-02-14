package main

func ValidLoginAttempt(email string, password string) (string, bool) {
	id, err := RepoGetUserIdByEmail(email)
	if err != nil {
		return id, false
	}
	hashed, err := RepoGetUserPasswordById(id)
	if err != nil {
		return id, false
	}
	ok := ComparePassword(hashed, password)
	return id, ok
}
