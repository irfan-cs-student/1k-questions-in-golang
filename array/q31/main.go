package main

import "fmt"

// Reverse an array without creating another array.

func reverseArray(a [6]int) [6]int {

	index := len(a) - 1
	center := len(a) / 2

	for i := 0; i < center; i++ {

		temp := a[i]
		a[i] = a[index]
		a[index] = temp
		index--

	}

	return a
}

func main() {

	nums := [6]int{2, 1, 7, 0, 4, -2}

	fmt.Print("pairs array:", reverseArray(nums))
}
