// Second-largest value
// Find the second-largest distinct number in a slice.
// Example:
// [10 50 20 50 30]
// → 30
// Don't sort the slice.

package main

import (
	"fmt"
)

func secondLarg(a []int) int {

	if len(a) == 0 {
		return -1
	}
	var b []int

	for _, value := range a {
		found := false
		for _, element := range b {

			if element == value {
				found = true
				break
			}
		}
		if !found {

			b = append(b, value)

		}
	}
	big, secondBig := 0, 0
	for _, value := range b {
		if big < value {
			secondBig = big
			big = value
		} else if big > value && secondBig < value {
			secondBig = value
		}
	}
	return secondBig
}

func main() {
	nums := []int{10, 50, 20, 50, 30}

	fmt.Print("2nd largest:", secondLarg(nums))
}
