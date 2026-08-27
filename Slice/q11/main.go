// Remove duplicates
// Given:
// [1 2 2 3 4 4 5 1]
// produce:
// [1 2 3 4 5]
// Preserve the first occurrence order.

package main

import "fmt"

func removeOccur(a []int) []int {
	var b []int

	for _, value := range a {

		found := false

		for _, element := range b {
			if value == element {
				found = true
				break
			}
		}

		if !found {
			b = append(b, value)
		}
	}

	return b
}
func main() {
	nums := []int{1, 2, 2, 3, 4, 4, 5}
	updatedSlice := removeOccur(nums)
	fmt.Print("updated:", updatedSlice)
}
