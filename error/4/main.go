package main

import (
	"errors"
	"fmt"
)

func findUser(id int) (string, error) {

	if id == 1 {
		return "Irfan", nil
	}

	if id == 2 {
		return "Ali", nil
	}

	if id == 3 {
		return "Ahmed", nil
	}

	return "", errors.New("user not found")
}

func main() {

	name, err := findUser(2)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User:", name)
}
