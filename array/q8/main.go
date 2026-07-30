package main

import "fmt"

// Reverse an array

func main() {

	nums := [5]int{10, 20, 30, 40, 50}
	reverseNum := [5]int{}

	i := 0

	for a := len(nums) - 1; a >= 0; a-- {
		reverseNum[i] = nums[a]
		i++
	}

	fmt.Println(reverseNum)
}
