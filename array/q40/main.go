package main

import (
	"fmt"
)

// Find the longest sequence of consecutive equal values.
func consValue(arr [10]int) (int, int) {

	maxCount, elememt, count := 1, arr[0], 0

	for i := 1; i < len(arr); i++ {

		if arr[i] == arr[i-1] {
			count++
		} else {
			count = 1
		}
		if count > maxCount {
			elememt = arr[i]
			maxCount = count
		}

	}
	return elememt, maxCount

}
func main() {

	num := [10]int{1, 1, 2, 2, 2, 3, 3, 3, 3, 1}
	elememt, maxCount := consValue(num)

	fmt.Println("value/element :", elememt, "repeat:", maxCount)

}
