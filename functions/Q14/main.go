package main

import "fmt"

// Swap two numbers using multiple return values.

func swap(i, j int) (int, int) {
	k := i
	i = j
	j = k
	return i, j
}
func main() {
	var a, b int
	fmt.Print("enter 1st num: ")
	fmt.Scan(&a)
	fmt.Print("enter 2nd num: ")
	fmt.Scan(&b)

	fmt.Println("---------your value bfore swap-------------")
	fmt.Println("1st value:", a)
	fmt.Println("2nd value:", b)

	fmt.Println("---------your value After swap-------------")
	aa, bb := swap(a, b)
	fmt.Println("1st value:", aa)
	fmt.Println("2nd value:", bb)

}
