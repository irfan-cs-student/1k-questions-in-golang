package main

import "fmt"

// Write a function that receives a slice and
// returns both the smallest and largest values.

func small_largest(a []int) (int, int) {

	small := a[0]
	largest := a[0]

	for _, value := range a {

		if small > value {
			small = value
		}
		if largest < value {
			largest = value
		}
	}

	return small, largest
}

func main() {

	slice := []int{1, 2, 3, 4, 5, 6}

	smallest, largest := small_largest(slice)

	fmt.Println("smallest:", smallest)
	fmt.Println("largest:", largest)
}
