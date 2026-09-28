// Create a map named scores using make().

// Add 3 students with their marks.

// Update one student's marks.

// Print the final map.

package main

import (
	"fmt"
)

func main() {

	//map using literal
	scores := map[string]int{

		"ali":   2,
		"usman": 3,
	}
	fmt.Println(scores)
	fmt.Println()

	//map using make
	s := make(map[string]int)

	s["english"] = 20
	s["urdu"] = 25
	s["math"] = 26

	fmt.Println("total marks in subjects:", s)

	fmt.Println("math after updating: ", s["math"])
	//upsating math marks
	s["math"] = 10
	fmt.Println("math after updating: ", s["math"])

	// deleting math
	delete(s, "math")
	fmt.Println("math after updating: ", s["math"])
	fmt.Println("\n _____total marks in subjects:____", s)

}
