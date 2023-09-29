package helpers

import (
	"task-5-pbi-btpns-arufhakim/controllers/queries"

	"golang.org/x/crypto/bcrypt"
)

var user queries.UserQuery

func IsRegistered(email string) (bool, uint, error) {
	user, err := user.Get(email)

	if err != nil {
		return false, 0, err
	}

	if user.ID == 0 {
		return false, 0, nil
	}

	return true, user.ID, nil
}

func HashPassword(password string) (string, error) {
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(HashedPassword), nil
}

func ComparePassword(password, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}

	return nil
}
