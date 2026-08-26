// Create a slice of 8 integers. Print it, change the value at index 3, then print the modified slice.
package main

import "fmt"

func main() {

	a := []int{3, 4, 2, 4, 2, 4, 5, 5, 3}
	a[3] = 99
	fmt.Print(a)

}
