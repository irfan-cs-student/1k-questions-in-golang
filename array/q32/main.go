package main

import "fmt"

// Return bigest and smallest number and their diference

func extremeNum(a [6]int) (largest, smallest int) {

	largest = a[0]
	smallest = a[0]

	for _, value := range a {

		if value > largest {
			largest = value
		}
		if value < smallest {
			smallest = value
		}

	}

	return
}

func main() {

	nums := [6]int{2, 1, 7, 0, 4, -2}

	smallest, largest := extremeNum(nums)
	fmt.Print("diference:", largest, " - ", smallest, "=", largest-smallest)

}
