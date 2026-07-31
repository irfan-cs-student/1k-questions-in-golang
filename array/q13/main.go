package main

import "fmt"

//Find duplicate values

func average(a, b int) int {
	return a / b
}

func main() {

	nums := [5]int{10, 20, 30, 40, 50}

	sum := 0

	for _, value := range nums {

		sum += value
	}
	fmt.Println("averagr : ", average(sum, len(nums)))
}
