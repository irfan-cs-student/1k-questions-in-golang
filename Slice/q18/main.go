// Write a program that receives:

// nums := []int{5, 2, 8, 2, 9, 1, 5, 8, 3, 2}

// Your program must:

// Remove duplicates.
// Keep the original order.
// Find the smallest value.
// Find the largest value.
// Calculate the average.
// Reverse the resulting slice.
// Print the final slice.

package main

import (
	"fmt"
)

func doOpperations(a []int) {
	if len(a) == 0 {
		fmt.Print("slice is empty !!")
		return
	}
	var result []int

	//removing duplicates
	for _, value := range a {

		found := false
		for _, element := range result {
			if value == element {
				found = true
				break
			}
		}

		if !found {
			result = append(result, value)

		}

	}
	fmt.Println("slice with out duplicates:", result)

	//find largest and smallest
	smallest, largest := result[0], result[0]

	for _, value := range result {
		if largest < value {
			largest = value
		}
		if smallest > value {
			smallest = value
		}
	}
	fmt.Println("largest value:", largest)
	fmt.Println("smallest value:", smallest)

	//average of all elements
	sum := 0
	for _, value := range result {

		sum += value
	}
	average := float64(sum) / float64(len(result))

	fmt.Println("average of the slice elements:", average)

	//reversing the array

	left, right := 0, len(result)-1

	for left < right {
		result[left], result[right] = result[right], result[left]

		left++
		right--
	}
	fmt.Print("reverse slice :", result)

}
func main() {

	nums := []int{5, 2, 8, 2, 9, 1, 5, 8, 3, 2}
	fmt.Println("orignal slice:", nums)
	doOpperations(nums)

}
