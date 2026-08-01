package main

import "fmt"

//find smallest and its index

func smallestElement(a [8]int) (int, int) {

	smallest, index := a[0], 0

	for i := 0; i < len(a); i++ {

		if smallest > a[i] {

			smallest = a[i]
			index = i
		}

	}
	return smallest, index
}

func main() {

	nums := [8]int{5, 2, 7, 2, 9, -7, 4, 8}

	smallest, index := smallestElement(nums)

	fmt.Println("Smallest element:", smallest, "\n index of smallest  element", index)

}
