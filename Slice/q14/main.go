// Rotate a slice to the left by k positions.
// [1 2 3 4 5 6]
// k = 2
// → [3 4 5 6 1 2]
// Your code should also work when k > len(slice).

package main

import "fmt"

func updatedSlice(slice []int, k int) []int {

	if len(slice) == 0 {
		return slice
	}

	var b []int
	for i := k; i < len(slice); i++ {

		b = append(b, slice[i])
	}
	for i := 0; i < k; i++ {

		b = append(b, slice[i])

	}

	return b
}
func main() {
	nums := []int{2, 4, 6, 8, 22, 44}

	var move int
	fmt.Print("move elements by value of k:")
	fmt.Scan(&move)

	fmt.Println("orignal array:", nums)
	fmt.Print("updated slice:", updatedSlice(nums, move))

}
