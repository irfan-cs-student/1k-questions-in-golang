package main

import "fmt"

//Find the largest number

func main() {

	nums := [5]int{10, 77, 30, 40, 50}
	largest := nums[0]

	for _, value := range nums {

		if largest < value {
			largest = value
		}
	}
	fmt.Print("Largest: ", largest)
}
