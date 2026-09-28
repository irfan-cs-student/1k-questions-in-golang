// nested maps
//
//	Count Word Frequency
package main

import "fmt"

func main() {

	students := map[string]map[string]int{
		"class_1": {

			"irfan": 1,
			"ali":   2,
			"saeed": 3,
		},
		"class_2": {

			"yasir": 22,
			"noor":  33,
			"odh":   44,
		},
	}

	fmt.Println(students)
}
