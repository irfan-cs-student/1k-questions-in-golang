// Given a slice and a target number, count how many times the target appears.

// [2 5 2 8 2 9] target=2
// → 3

package main

import "fmt"

func frequency(a []int, b int) int {

	count := 0
	for _, value := range a {
		if b == value {
			count++
		}
	}
	return count
}
func main() {
	nums := []int{2, 4, 2, 3, 1, 3, 2, 3, 2}

	var a int
	fmt.Println("nums:", nums)
	fmt.Print("give target value :")
	fmt.Scan(&a)

	fmt.Println(a, "appears at frequency of :", frequency(nums, a))
}
