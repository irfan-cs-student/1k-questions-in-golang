package main

import "fmt"

//Find the maximum consecutive sum.
func maxConsecutiveSum(arr [6]int) int {
	maxSum, currentMax := arr[0], arr[0]

	for index, _ := range arr {

		i := index + 1
		if i >= len(arr) {
			break
		}

		if currentMax+arr[i] > arr[i] {

			currentMax = currentMax + arr[i]

		} else {
			currentMax = arr[i]
		}

		if currentMax > maxSum {
			maxSum = currentMax
		}

	}
	return maxSum
}
func main() {

	a := [6]int{4, -6, 3, 5, -2, 4}

	fmt.Println("consective sum:", maxConsecutiveSum(a))

}
