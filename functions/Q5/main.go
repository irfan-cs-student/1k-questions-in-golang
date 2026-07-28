package main

import "fmt"

// Write a function to print the multiplication table of a number.

func printTable(n int) {
	for a := 1; a <= 20; a++ {
		fmt.Println(a, " * ", n, " = ", a*n)
	}
}
func main() {
	var n int
	fmt.Scan(&n)

	fmt.Println("Table of :", n)
	printTable(n)
}
