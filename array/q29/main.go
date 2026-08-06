package main

import "fmt"

// Find two numbers whose sum equals a target

func pairOfsum(a [6]int, target int) (result [2]int, firstIndex, lastIndex int) {

	for i := 0; i < len(a); i++ {

		for j := i + 1; j < len(a); j++ {

			if a[i]+a[j] == target {

				result[0] = a[i]
				result[1] = a[j]
				firstIndex = i
				lastIndex = j
				return

			}

		}

	}
	return
}

func main() {

	nums := [6]int{23, 2, 3, 7, 2}
	target := 30

	pair, firstIndex, lastIndex := pairOfsum(nums, target)

	fmt.Println("pair: ", pair, "---", pair[0], "+", pair[1], "=", pair[0]+pair[1])
	fmt.Println("their index --------- \n first elemt index in main array :", firstIndex)
	fmt.Print(" second elemt index in main array :", lastIndex)

}
