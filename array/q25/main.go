package main

import (
	"fmt"
)

// Find the sum of only unique values
// Find the sum of distinct values

func sumOfUnique(abc [7]int) (int, int) {
	sum := 0
	distinctSum := 0

	for a := 0; a < len(abc); a++ {

		count := 0

		for b := 0; b < len(abc); b++ {

			if abc[a] == abc[b] {
				count++
			}
		}
		if count == 1 {
			sum += abc[a]
		}
	}

	for a := 0; a < len(abc); a++ {

		isDistinct := true

		for b := 0; b < a-1; b++ {

			if abc[a] == abc[b] {
				isDistinct = false
				break
			}
		}
		if isDistinct {
			distinctSum += abc[a]
		}
	}
	return sum, distinctSum
}
func main() {

	nums := [7]int{2, 3, 2, 5, 3, 7, 5}
	nique, distinct := sumOfUnique(nums)
	fmt.Println("sum of unique: ", nique)
	fmt.Println("sum of distinct: ", distinct)
}
