package main

import (
	"fmt"
)

func main() {
	pass1 := "12345"
	hashedPass, err := hashPassword(pass1)
	if err != nil {
		panic(err)
	}

	pass2 := "123453"
	err = comparePasswords(pass2, hashedPass)
	if err != nil {
		panic(err)
	}

	fmt.Println("passwords are equal")
}
