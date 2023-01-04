package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(pass string) (hashedPass []byte, err error) {
	hashedPass, err = bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash pass: %w", err)
	}

	return hashedPass, nil
}

func comparePasswords(pass string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(pass))
	if err != nil {
		return fmt.Errorf("passwords are not equal: %w", err)
	}

	return nil
}
