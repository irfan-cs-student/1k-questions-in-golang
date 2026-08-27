package main

import "fmt"

// Write a function that removes an element at a given index.

// [10 20 30 40 50]
// remove index 2
// → [10 20 40 50]

func deleteIndex(a []int, target int) []int {

	if target < 0 || target >= len(a) {
		return a
	}
	for i := target; i < len(a)-1; i++ {

		a[i] = a[i+1]
	}
	return a[:len(a)-1]
}
func main() {
	nums := []int{10, 20, 30, 40, 50}

	fmt.Println("original array:", nums)
	var target int
	fmt.Print("target value: ")
	fmt.Scan(&target)

	fmt.Print("updated array:", deleteIndex(nums, target))

}
