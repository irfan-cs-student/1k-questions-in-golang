package main

import "fmt"

// Write a function to print numbers from 1 to N.
func printNumbers(n int) {
	for a := 1; a <= n; a++ {
		fmt.Print(a, " ")
	}
}
func main() {
	var n int
	fmt.Print("enter number n: ")
	fmt.Scan(&n)
	printNumbers(n)
}
