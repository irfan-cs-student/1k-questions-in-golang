package main

import "fmt"

func alpha() {
	defer fmt.Println("A1")
	fmt.Println("A2")
	panic("STOP")
	fmt.Println("A3")
}

func beta() {
	defer fmt.Println("B1")
	alpha()
	fmt.Println("B2")
}

func main() {
	defer fmt.Println("M1")
	beta()
	fmt.Println("M2")
}

//output:
/*
main put defer in lifo stack
main() calls beta
beta() put defer in stack "b1"
 and calls alpha()
 {
 alpha puts defer into stack lifo "a1"
 prints a2
 panic rewinds stack and execute a1
 so ------a2,a1
 }
 and then prints "m1"
 and later m2 after alpha calls

 result:
 a2,a1
 b1
 m1

*/
