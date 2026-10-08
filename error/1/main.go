package main

import "fmt"

// Marks Validation

// Create:

// func checkMarks(marks int) error

// Rules:

// Below 0 → "invalid marks"
// Above 100 → "invalid marks"
// Otherwise → nil

// Test with:
type Student struct {
	name  string
	marks int
}

// checkMarks(120)
func checkMarks(s Student) error {

	if s.marks > 100 || s.marks < 0 {
		return fmt.Errorf("%s student's marks is invalid", s.name)
	}
	return nil
}
func main() {

	students := []Student{
		{name: "Ali", marks: 75},
		{name: "Ahmed", marks: -82},
		{name: "Usman", marks: 68},
		{name: "Hassan", marks: 91},
		{name: "Bilal", marks: -77},
		{name: "Hamza", marks: 101},
		{name: "Talha", marks: -64},
		{name: "Saad", marks: 95},
		{name: "Zain", marks: 73},
		{name: "Umar", marks: 86},
	}

	for i := range students {
		s := checkMarks(students[i])

		if s != nil {
			fmt.Println(s)
		} else {
			fmt.Println(students[i].name, "has valid marks")

		}
	}

}
