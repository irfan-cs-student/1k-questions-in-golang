package main

import "fmt"

//find largest and its index

func largeElement(a [8]int) (int, int) {

	largest, index := a[0], 0

	for i := 0; i < len(a); i++ {

		if largest < a[i] {

			largest = a[i]
			index = i
		}

	}
	return largest, index
}

func main() {

	nums := [8]int{5, 2, 7, 2, 9, 2, 4, 8}

	largest, largest_index := largeElement(nums)

	fmt.Println("largest element:", largest, "\n index of its largest element", largest_index)

}
