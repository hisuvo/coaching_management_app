package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Task :  password hashing and comparing task

func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password is required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// compare password
func ComparePassword(hash string, password string) error {

	err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(password),
	)

	return err
}