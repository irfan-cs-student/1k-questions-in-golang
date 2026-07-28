package main

import "fmt"

// Write a function to calculate the sum of numbers from 1 to N.

func printSum(n int) int {
	sum := 0
	for a := 1; a <= n; a++ {
		sum += a
	}
	return sum
}
func main() {
	var n int
	fmt.Print("enter number n: ")
	fmt.Scan(&n)
	fmt.Println("sum from ", 1, "to", n, "=", printSum(n))
}
