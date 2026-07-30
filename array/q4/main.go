package main

import "fmt"

//Find the smallest  number

func main() {

	nums := [5]int{67, 77, 30, 40, 50}
	smallest := nums[0]

	for _, value := range nums {

		if smallest > value {
			smallest = value
		}
	}
	fmt.Print("smalllest: ", smallest)
}
