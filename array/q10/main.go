package main

import "fmt"

// Count a specific value
func countNumber(nums [7]int, target int) int {

	count := 0
	for _, value := range nums {

		if value == target {
			count++
		}
	}
	return count

}

func main() {

	nums := [7]int{10, 20, 30, 40, 50, 20, 20}
	a := 20

	fmt.Println("specialNum repeated times: ", countNumber(nums, a))

}
