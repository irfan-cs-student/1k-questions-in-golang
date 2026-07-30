package main

import "fmt"

// Print all elements

func main() {

	nums := [5]int{10, 20, 30, 40, 50}

	for _, value := range nums {

		fmt.Print(value)

	}
}
