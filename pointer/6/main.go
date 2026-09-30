package main

import "fmt"

func change(p *int) {
	*p = 100 //p has adress ,and *p=a --> a=100
}

func main() {
	a := 20

	change(&a)

	fmt.Println(a) //100
}
