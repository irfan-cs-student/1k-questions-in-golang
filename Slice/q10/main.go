package main

import "fmt"

// Insert a value at a specific index.

// [10 20 40 50]
// insert 30 at index 2
// → [10 20 30 40 50]
// Don't use a library/helper that directly performs insertion.

func deleteIndex(a []int, target int, value int) []int {

	if target < 0 || target >= len(a) {
		return a
	}
	a = append(a, 0)

	for i := len(a) - 1; i > target; i-- {

		a[i] = a[i-1]
	}
	a[target] = value

	return a
}
func main() {
	nums := []int{10, 20, 30, 40, 50}

	fmt.Println("original array:", nums)
	var target int
	fmt.Print("target index: ")
	fmt.Scan(&target)
	var value int
	fmt.Print("value to store:")
	fmt.Scan(&value)

	fmt.Print("updated array:", deleteIndex(nums, target, value))

}
