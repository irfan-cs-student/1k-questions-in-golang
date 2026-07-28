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
func checkPalindrome(a, b int) {
	if a == b {
		fmt.Println(a, "=", b, "---so ,number is palidrome")
	} else {
		fmt.Println(a, "not equal to ", b, "-----so, number is not palidrome")

	}
}
func main() {
	var n int
	fmt.Print("enter number n: ")
	fmt.Scan(&n)
	reverse_number := printReverse(n)
	checkPalindrome(n, reverse_number)

}
