package main

import (
	"errors"
	"fmt"
)

type AgeError struct {
	Age int
}

func (e AgeError) Error() string {
	return fmt.Sprintf("invalid age: %d", e.Age)
}

func checkAge(age int) error {
	if age < 18 {
		return AgeError{Age: age}
	}

	return nil
}

func main() {
	err := checkAge(15)

	var ageErr AgeError

	if errors.As(err, &ageErr) {
		fmt.Println("Age error found")
		fmt.Println("Age:", ageErr.Age)
	}
}
