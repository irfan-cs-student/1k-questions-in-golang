package main

import "fmt"

//count of elememts more than 50 values

func count(a [8]int) int {

	count := 0

	for i := 0; i < len(a); i++ {

		if a[i] > 50 {

			count++

		}

	}
	return count
}

func main() {
	nums := [8]int{12, 67, 34, 89, 51, 23, 90, 45}

	fmt.Println("count of elememts more than 50 values:", count(nums))

}
