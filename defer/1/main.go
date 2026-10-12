package main

import "fmt"

func test() {
	defer fmt.Println("A")
	fmt.Println("B")
}

func main() {
	test()

	//output
	//b
	//a
}
