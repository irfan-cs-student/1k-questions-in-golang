// nested maps
//
//	Count Word Frequency
package main

import (
	"fmt"
)

func main() {

	students := map[string]map[string]int{
		"prep": {

			"irfan": 1,
			"ali":   2,
			"saeed": 3,
		},
		"one": {

			"yasir": 22,
			"noor":  33,
			"odh":   44,
		},
	}

	fmt.Println()

	for class, studentmap := range students {

		fmt.Println("class:_______", class)
		fmt.Println()

		for name, value := range studentmap {

			fmt.Println(name, "--", value)

		}
		fmt.Println()

	}
}
