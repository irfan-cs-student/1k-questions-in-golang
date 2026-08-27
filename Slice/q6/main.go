package main

import "fmt"

// Reverse in-place
// Reverse a slice without creating another slice.
// Example:
// [1 2 3 4 5] → [5 4 3 2 1]

func Reverse(a []int) []int {

	left := 0
	right := len(a) - 1

	for left < right {

		a[left], a[right] = a[right], a[left]

		left++
		right--
	}

	return a
}

func main() {

	slice := []int{1, 2, 3, 4, 5, 6}
	Reverse := Reverse(slice)
	fmt.Print(Reverse)
}
