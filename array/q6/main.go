package main

import "fmt"

//Count positive and negative numbers

func main() {

	nums := [...]int{67, 77, -30, 40, 50, -47, 88, -99, 21}
	posCount := 0
	negCount := 0

	for _, value := range nums {

		if value >= 0 {
			posCount++

		} else {
			negCount++
		}
	}
	fmt.Println("positive numbers: ", posCount)
	fmt.Print("negative numbers: ", negCount)
}
