package main

import (
	"fmt"
)

//Find the first and last occurrence
//Find the first and last index of 2.

func occure(value int, a [8]int) (int, int) {

	first, last := -1, -1

	for i := 0; i < len(a); i++ {

		if a[i] == value {

			if first == -1 {

				first = i
			}
			last = i

		}

	}
	return first, last
}

func main() {

	nums := [8]int{5, 2, 7, 2, 9, 2, 4, 8}

	value := 0
	fmt.Scan(&value)

	first, last := occure(value, nums)
	fmt.Println("first occurence: ", first, " \n last occurrence: ", last)

	if first == last {
		fmt.Print("first is also last")
	}

}
