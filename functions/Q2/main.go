package main

import "fmt"

//Write a function to print numbers from N to 1.

func printNumbers(n int) {

	if n >= 1 {

		for a := n; a >= 1; a-- {
			fmt.Print(a, " ")
		}
	} else {
		for a := n; a <= 1; a++ {
			fmt.Print(a, " ")
		}
	}
}
func main() {
	var n int
	fmt.Print("enter number n: ")
	fmt.Scan(&n)
	printNumbers(n)
}
