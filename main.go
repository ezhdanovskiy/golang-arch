package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass1 := "12345"
	hashedPass, err := hashPass(pass1)
	if err != nil {
		panic(err)
	}

	pass2 := "123453"
	err = compare(pass2, hashedPass)
	if err != nil {
		panic(err)
	}

	fmt.Println("passwords are equal")
}

func hashPass(pass string) (hashedPass []byte, err error) {
	hashedPass, err = bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash pass: %w", err)
	}

	return hashedPass, nil
}

func compare(pass string, hashedPass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPass, []byte(pass))
	if err != nil {
		return fmt.Errorf("passwords are not equal: %w", err)
	}

	return nil
}
