package main

import "fmt"

// Move all zeros to the end

func changeArray(a [9]int) [9]int {

	// Traverse remaining elements
	var b [9]int
	index := 0

	for i := 0; i < len(a); i++ {

		if a[i] == 0 {

			continue
		}
		b[index] = a[i]
		index++

	}

	return b
}

func main() {

	nums := [9]int{10, 0, -5, 0, 20, 0, -8, 30, -2}

	fmt.Println(changeArray(nums))
}
