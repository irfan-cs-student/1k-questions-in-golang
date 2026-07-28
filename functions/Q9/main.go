package main

import "fmt"

//Write a function to reverse a number.
func printReverse(n int) int {
	last_diget, reverse_num := 0, 0

	for n != 0 {
		last_diget = n % 10
		reverse_num = reverse_num*10 + last_diget
		n = n / 10

	}
	return reverse_num
}
func main() {
	var n int
	fmt.Print("enter number n: ")
	fmt.Scan(&n)
	fmt.Println("reverse number=", printReverse(n))
}
