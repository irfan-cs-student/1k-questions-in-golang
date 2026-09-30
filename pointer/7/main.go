package main

import "fmt"

func main() {
	a := 10
	p := &a //*p=10

	fmt.Println(p)  //adress of a
	fmt.Println(*p) //*p=a=10
	fmt.Println(&p) //&p=adrees of pointer p
}
