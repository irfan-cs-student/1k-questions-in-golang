package main

import "fmt"

// Write a function to calculate the factorial of a number.

func printFact(n int) int {
	fact := 1
	for a := 1; a <= n; a++ {
		fact *= a
	}
	return fact
}
func main() {
	var n int
	fmt.Print("enter number n: ")
	fmt.Scan(&n)
	fmt.Println("factorial of: ", n)

	if n < 0 {
		fmt.Println("factorail not defined for 'n<0'case ")
	} else if n == 1 || n == 0 {
		fmt.Println("factorial is =", 1)
	} else {

		fmt.Println("factorial is =", printFact(n))
	}
}
