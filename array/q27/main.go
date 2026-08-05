package main

import "fmt"

// array is palindrome or not
func Ispalindrome(a [5]int) bool {

	var b [5]int
	indexB := len(a) - 1

	for _, value := range a {

		b[indexB] = value
		indexB--
	}
	//chekig reverse and orignal array values same or diferenet
	index2B := 0
	IsSame := true

	for _, value := range a {

		if value != b[index2B] {
			IsSame = false
			break
		}
		index2B++

	}
	return IsSame
}

func main() {

	nums := [5]int{1, 2, 3, 9, 1}

	fmt.Print("palindrome :", Ispalindrome(nums))

}
