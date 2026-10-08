package main

//error.Is()

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("user not found")
var ErrBlocked = errors.New("user is blocked")

func checkUser(id int) error {
	if id == 1 {
		return ErrNotFound
	}

	if id == 2 {
		return ErrBlocked
	}

	return nil
}

func main() {
	err := checkUser(2)

	if errors.Is(err, ErrNotFound) {
		fmt.Println("User does not exist")
	}

	if errors.Is(err, ErrBlocked) {
		//its chek is err and errbloked both same or diferent
		//  and allow condintion true or false and as allowing printing or not
		fmt.Println("User is blocked")
	}
}
