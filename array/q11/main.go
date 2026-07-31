package main

import "fmt"

// Find the second largest number
func largestNum(nums [7]int) (int, int) {

	largest, _2ndLargest := 0, 0

	for _, value := range nums {

		if value > largest {
			_2ndLargest = largest
			largest = value

		} else if value > _2ndLargest {
			_2ndLargest = value
		}
	}
	return largest, _2ndLargest

}

func main() {

	nums := [7]int{10, 20, 80, 48, 50, 57, 29}

	largest, _2ndLargest := largestNum(nums)
	fmt.Println("largest number: ", largest, "2nd_laragest number: ", _2ndLargest)

}
