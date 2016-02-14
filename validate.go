package main

func ValidLoginAttempt(email string, password string) (string, bool, error) {
	id, err := RepoGetUserIdByEmail(email)
	if err != nil {
		return id, false, err
	}
	hashed, err := RepoGetUserPasswordById(id)
	if err != nil {
		return id, false, err
	}
	ok := ComparePassword(hashed, password)
	return id, ok, nil
}
