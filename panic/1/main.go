package main

import "fmt"

func test() {
	defer fmt.Println("1")
	fmt.Println("2")

	panic("Oops")

	defer fmt.Println("3")
	fmt.Println("4")
}

func main() {
	fmt.Println("Start")
	test()
	fmt.Println("End")
}

//otput-------
//start
//2
//1
//oops ,bacuse first panic rewinds stack and then shows panic
