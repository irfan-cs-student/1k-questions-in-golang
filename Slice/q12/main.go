// Given a slice of integers, create two slices:

// even → all even numbers
// odd  → all odd numbers

// Preserve their original order.

package main

import (
	"fmt"
)

func evenOdd(a []int) (b []int, c []int) {

	if len(a) == 0 {
		return
	}
	for _, value := range a {

		if value%2 == 0 {

			b = append(b, value)

		} else {
			c = append(c, value)
		}
	}
	return
}
func main() {
	nums := []int{}
	// {2, 5, 23, 5, 3, 4, 5, 6, 8, 10, 22, 34, 2, 1}

	even, odd := evenOdd(nums)
	fmt.Print("even:", even, "\nodd:", odd)
}
