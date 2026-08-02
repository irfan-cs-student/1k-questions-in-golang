package main

import "fmt"

// Find missing number from 1 to 9

func changeArray(a [9]int) [9]int {

	// Traverse remaining elements
	var b [9]int
	index := 0

	for i := 0; i < len(a); i++ {

		found := false

		for _, value := range a {

			if value == i {
				found = true
				break
			}
		}
		if !found {
			b[index] = i
			index++
		}

	}

	return b
}

func main() {

	nums := [9]int{1, 3, 5, 7, 8}

	// fmt.Println(changeArray(nums))
	fmt.Print(changeArray(nums))
}
