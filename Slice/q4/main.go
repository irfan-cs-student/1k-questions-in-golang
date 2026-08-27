package main

import "fmt"

// Write a function that receives a slice of integers and
// returns its first and last elements. Handle an empty slice safely.

func getElements(a []int) (int, int) {

	last_index := len(a) - 1
	return a[0], a[last_index]

}

func main() {

	nums := []int{4, 5, 6, 2, 5, 24, 5, 3}

	_1st, lastDigit := getElements(nums)

	fmt.Print("first:", _1st, "\nlast:", lastDigit)

}
