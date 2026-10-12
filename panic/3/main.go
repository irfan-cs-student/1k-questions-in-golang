package main

import "fmt"

func test() {
	defer fmt.Println("A")

	defer func() {
		fmt.Println("B")
		panic("SECOND")
	}()

	defer fmt.Println("C")

	panic("FIRST")
}

func main() {
	defer fmt.Println("D")
	test()
}

//output_____
/*
a
b
d
second
*/
