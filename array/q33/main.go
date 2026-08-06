package main

import "fmt"

// Return bigest and second bigest number and their diference

func extremeNum(a [6]int) (largest, secLargest int) {

	if a[0] > a[1] {
		largest = a[0]
		secLargest = a[1]
	} else {
		largest = a[1]
		secLargest = a[0]
	}

	for _, value := range a {

		if value > largest {
			secLargest = largest
			largest = value
		} else if value < largest && value > secLargest {

			secLargest = value
		}

	}

	return
}

func main() {

	nums := [6]int{2, 1, 7, 0, 4, -2}

	largest, secondLarg := extremeNum(nums)
	fmt.Print("diference:", largest, " - ", secondLarg, "=", largest-secondLarg)

}
