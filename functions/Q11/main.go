package main

import "fmt"

//calling all functions through single functions
func add(a, b int) {
	fmt.Println(a, "+", b, "=", a+b)
}
func sub(a, b int) {
	fmt.Println(a, "-", b, "=", a-b)
}
func multiply(a, b int) {
	fmt.Println(a, "*", b, "=", a*b)
}
func divide(a, b int) {
	fmt.Println(a, "/", b, "=", a/b)
}
func allCalculations(a, b int) {

	add(a, b)
	sub(a, b)
	multiply(a, b)
	divide(a, b)

}
func main() {
	var a, b int
	fmt.Print("give 1st num:")
	fmt.Scan(&a)
	fmt.Print("give 2nd num:")
	fmt.Scan(&b)

	fmt.Println("-------Basic arithmatic calculations----")
	allCalculations(a, b)
}
