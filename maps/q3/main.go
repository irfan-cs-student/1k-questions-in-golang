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

	//ok concept
	math, ok := s["math"]
	if ok {
		fmt.Println("math marks:", math)

	} else {
		fmt.Println("math subject not exist !")

	}

	// deleting math
	delete(s, "math")
	fmt.Println("math after updating: ", s["math"])

	//ok concept
	math_marks, ok := s["math"]

	if ok {
		fmt.Println("math marks:", math_marks)

	} else {
		fmt.Println("math subject not exist !")

	}

	fmt.Println("\n _____total marks in subjects:____", s)

	fmt.Println("\n \n \n  \n  \n  \n \n")

	//-----------------------
	//maps creations with diferent syntax

	var m map[string]int //nill map
	// m["ali"] = 2 causr painc becuse we cant add in nill map
	fmt.Println(m == nil) //true

	// for insertion in nill map
	m = make(map[string]int)
	m["ali"] = 2
	fmt.Println(m)

	var n = make(map[string]int)
	n["irfan"] = 9
	fmt.Println(n)

	var o = map[int]string{
		1: "irfan",
	}
	fmt.Println(o)

}
