// Sum without modifying
// Write a function:

// sum(nums []int) int
// that returns the sum of all elements.

package main

import "fmt"

func sumOfelements(a []int) int {

	sum := 0
	for _, value := range a {

		sum += value
	}
	return sum
}
func main() {
	nums := []int{10, 20, 30, 40, 50, 60, 70, 80}

	sum := sumOfelements(nums)
	fmt.Print(sum)

}
