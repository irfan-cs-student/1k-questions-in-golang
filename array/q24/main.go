package main

import "fmt"

func findSmall(a [6]int) (int, int) {

	// first two elements

	var smallest, secondSmallest int
	var secondIndex int

	if a[0] < a[1] {
		smallest = a[0]
		secondSmallest = a[1]
		secondIndex = 1
	} else {
		smallest = a[1]
		secondSmallest = a[0]
		secondIndex = 0
	}

	// Traverse remaining elements
	for i := 2; i < len(a); i++ {

		if a[i] < smallest {
			secondSmallest = smallest
			smallest = a[i]

		} else if a[i] > smallest && a[i] < secondSmallest {
			secondSmallest = a[i]
			secondIndex = i
		}
	}

	return secondSmallest, secondIndex
}

func main() {

	nums := [6]int{10, -5, 20, -8, 30, -2}

	secondSmallest, index := findSmall(nums)

	fmt.Println("Second Smallest:", secondSmallest)
	fmt.Println("Index:", index)
}
