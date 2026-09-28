// Create a map of 5 students and their marks.

// Find the student with the highest marks.

// Print the student's name and marks.
package main

import "fmt"

func main() {
	scores := map[string]int{
		"Ali":   85,
		"Ahmed": 92,
		"Sara":  95,
		"Usman": 78,
		"Bilal": 88,
	}
	highest := 0
	nameOftopper := ""

	for name, value := range scores {

		if value > highest {
			highest = value
			nameOftopper = name
		}
	}

	fmt.Println("name of topper:", nameOftopper, "marks:", highest)

}
