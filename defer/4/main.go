package main

import "fmt"

func test() int {
	defer fmt.Println("Finished")
	return 25
}

func main() {
	x := test()
	fmt.Println("Value:", x)
}

//main calls test()
//test execute and returns 25
//  and defer execute after returns as its stores as in stack as lifo
//return values assign to x and then print
