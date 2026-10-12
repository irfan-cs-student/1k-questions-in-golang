package main

import "fmt"

func test() {
	x := 10
	defer fmt.Println(x)
	//x gets imedate value 10 as argument but later execute

	x = 20
	fmt.Println(x) //prints 20
}

func main() {
	test()
}

//output
//20
//10
