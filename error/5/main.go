package main

import (
	"errors"
	"fmt"
)

func findUser(id int) error {
	if id != 10 {
		return errors.New("user not found")
	}

	return nil
}

func getUser(id int) error {
	err := findUser(id)

	if err != nil {
		return fmt.Errorf("failed to get user %d: %w", id, err)
	}

	return nil
}

func main() {
	err := getUser(5)

	if err != nil {
		fmt.Println(err)
	}
}
