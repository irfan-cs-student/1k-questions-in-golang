// Task:

// Add "Irfan": 95
// Change Ahmed's score to 100
// Delete Sara
// Loop through the map
// Find the highest score and student

package main

import "fmt"

func main() {

	students := map[string]int{
		"Ali":   80,
		"Ahmed": 90,
		"Sara":  75,
	}

	for name, value := range students {
		fmt.Println(name, " -- ", value)
	}

	students["irfan"] = 95
	students["Ahmed"] = 100
	delete(students, "Sara")

	_, ok := students["Sara"]

	if ok {
		fmt.Print("sara exist")
	} else {
		fmt.Print("sara deleted !!")

	}

	fmt.Println("_____________after operations____")
	for name, value := range students {
		fmt.Println(name, " -- ", value)
	}

	highest := 0
	name_of_topper := ""

	for name, value := range students {

		if highest < value {
			highest = value
			name_of_topper = name
		}
	}

	fmt.Println("name of topper:", name_of_topper, " marks:", highest)

}
