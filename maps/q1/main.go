// 1. Student Score Analyzer

// Create a program that stores student names and their scores in a map[string][]int.

// For each student:

// Calculate their average score.
// Find their highest score.
// Find their lowest score.
// Print students whose average is ≥ 80.

package main

import "fmt"

func main() {
	students := map[string][]int{

		"Ali":   {85, 90, 78, 92},
		"Ahmed": {70, 75, 80, 72},
		"Sara":  {95, 88, 92, 96},
		"Usman": {60, 65, 70, 68},
	}

	for name, value := range students {

		sum := 0

		highest := value[0]
		lowest := value[0]

		for _, score := range value {

			sum += score

			if highest < score {
				highest = score
			} else if lowest > score {
				lowest = score

			}

		}
		average := sum / len(value)

		fmt.Println("name :", name)
		fmt.Println("average :", average)
		fmt.Println("highest :", highest)
		fmt.Println("lowest :", lowest)

		if average >= 80 {
			fmt.Println(name, "has average greather than or equall to 80")
		}
		fmt.Println()
	}

}
