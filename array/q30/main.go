package main

import "fmt"

// Find all pairs whose sum equals a target.

func pairOfsum(a [6]int, target int) (result [6]int) {

	index := 0
	for i := 0; i < len(a); i++ {

		for j := i + 1; j < len(a); j++ {

			if a[i]+a[j] == target {

				result[index] = a[i]
				index++
				result[index] = a[j]
				index++
			}
		}
	}
	return
}

func main() {

	nums := [6]int{2, 1, 7, 0, 4, -2}
	target := 5

	fmt.Print("pairs array:", pairOfsum(nums, target))
}
