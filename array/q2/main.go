package main

import "fmt"

//Find the sum

func main() {

	nums := [5]int{10, 20, 30, 40, 50}
	sum := 0

	for _, value := range nums {

		sum += value

	}
	fmt.Print("sum:", sum)
}
