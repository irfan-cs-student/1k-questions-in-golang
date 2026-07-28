package main

import "fmt"

// Return the first and last digit of a number.

func first_last_digit(n int) (int, int) {
	last_digit := n % 10

	var first_digit int
	for n >= 10 {
		n /= 10
		first_digit = n

	}

	return first_digit, last_digit
}
func main() {
	var n int
	fmt.Print("enter number n: ")
	fmt.Scan(&n)
	_1st, _nth := first_last_digit(n)
	fmt.Println("first digit:", _1st)
	fmt.Println("Last digit:", _nth)

}
