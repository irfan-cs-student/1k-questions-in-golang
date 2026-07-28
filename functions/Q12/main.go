package main

import "fmt"

//Return quotient and remainder.

func printNumbers(a, b int) (int, int) {
	return a / b, a % b
}
func main() {
	var a, b int
	fmt.Print("enter number a and b : ")
	fmt.Scan(&a, &b)
	quotient, remainder := printNumbers(a, b)

	fmt.Println("qoutient:", quotient, "remainder:", remainder)
}
