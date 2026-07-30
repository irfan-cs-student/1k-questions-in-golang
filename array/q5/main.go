package main

import "fmt"

//Count even numbers
//Count odd numbers

func main() {

	nums := [...]int{67, 77, 30, 40, 50, 47, 88, 99, 21}
	Even_count := 0
	Odd_count := 0

	for _, value := range nums {

		if value%2 == 0 {
			Even_count++

		} else {
			Odd_count++
		}
	}
	fmt.Println("Even numbers ", Even_count)
	fmt.Print("Odd numbers ", Odd_count)
}
