package main

import "fmt"

func test() int {
	defer fmt.Println("A")
	defer fmt.Println("B")
	return 50
}

func main() {
	fmt.Println(test())
}

// output
//B
//A
//50
