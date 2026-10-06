package main

import "fmt"

// 3. Create a custom type:
// type Number int
// Write a method double() that returns twice the number.
type Number int

func (a Number) twice() Number {

	return a * a
}
func main() {

	n := Number(3)
	x := n.twice()
	fmt.Print(x)
}
