// Create a function modify() that accepts three integer pointers.

// Tasks:

// Add 10 to the first variable.
// Multiply the second variable by the updated first variable.
// Set the third variable equal to the sum of the first and second variables.
// Swap the values of the first and third variables using pointers.

// Starting values:

// a := 5
// b := 3
// c := 20

package main

import "fmt"

func modify(a, b, c *int) {

	fmt.Println("a: ", *a, " --b:", *b, "--c:", *c)

	*a = *a + 10 //a=15
	*b = *a * *b //15 *3=45
	*c = *a + *b // 15 + 45=60

	//swap
	temp := *a
	*a = *c
	*c = temp

	fmt.Println("a: ", *a, " --b:", *b, "--c:", *c)

}
func main() {

	a := 5
	b := 3
	c := 20

	modify(&a, &b, &c)

}
