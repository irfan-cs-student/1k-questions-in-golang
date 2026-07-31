package main

import "fmt"

// Check if array is sorted
func isSorted(a [5]int, b [7]int) (bool, bool) {

	sortA := true
	sortB := true
	for i := 1; i < len(a); i++ {

		if a[i] < a[i-1] {

			sortA = false
			break

		}

	}
	for i := 1; i < len(b); i++ {

		if b[i] < b[i-1] {

			sortB = false
			break
		}

	}

	return sortA, sortB

}

func main() {

	numsA := [5]int{10, 20, 30, 40, 50}
	numsB := [7]int{23, 45, 212, 3, 32, 92, 21}

	sortA, sortB := isSorted(numsA, numsB)

	fmt.Println("first array : ", numsA, "----sorted =", sortA)
	fmt.Println("2nd array: ", numsB, "-----sorted =", sortB)

}
