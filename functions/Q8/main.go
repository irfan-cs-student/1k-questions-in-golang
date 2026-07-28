package main

import "fmt"

//Write a function to count digits in a number.

func numCount(n int) int {
	count := 0
	if n == 0 {
		return 1
	}
	for n != 0 {
		n /= 10
		count++

	}
	return count

}
func main() {
	var n int
	fmt.Print("enter number: ")
	fmt.Scan(&n)
	fmt.Println("total diget in number :", numCount(n))

}
