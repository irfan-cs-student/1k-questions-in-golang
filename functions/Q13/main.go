// Return the larger and smaller number.
package main

import "fmt"

//Return quotient and remainder.

func printbigger(a, b int) (int, int) {
	if a > b {
		return a, b
	}
	return b, a

}
func main() {
	var a, b int
	fmt.Print("enter number a and b : ")
	fmt.Scan(&a, &b)

	bigger, smaller := printbigger(a, b)
	fmt.Println("biger:", bigger, "Smaller:", smaller)
}
