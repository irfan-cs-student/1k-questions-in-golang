package main

import (
	"errors"
	"fmt"
)

// user registration
func registerUser(username string, age int) error {

	if username == "" || age < 18 {
		if username == "" && age < 18 {
			return errors.New("invalid username && below 18 not allowed")
		}
		if username == "" {
			return errors.New("invalid username")
		} else if age < 18 {
			return errors.New("age is invalid")
		}
	}
	return nil
}

func main() {

	err := registerUser("", 9)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("users successfuly login")
	}
}
