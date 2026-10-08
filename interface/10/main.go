package main

import "fmt"

//empty interfaces & assertions
type Animal struct{ name string }

func main() {

	var x any = 3

	//var x interface{} = 3
	//var x any
	//then x=3 allowaed
	// x(any):=3 this not allowed syntx error

	fmt.Println(x)

	x = "irfan"
	fmt.Println(x)

	x = true
	fmt.Println(x)

	x = Animal{"goat"}
	fmt.Println(x)

	//n := x.(int) ----panic because n is int and x is struct

	n, ok := x.(int)
	fmt.Println(n)
	fmt.Println(ok)

}
