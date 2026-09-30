package main

import "fmt"

func updateValue(p **int) {
	x := 50
	*p = &x //where p pointing,now p changes direction to x so now **p is 50
}

func main() {
	a := 10
	// b := 20
	p := &a

	updateValue(&p) //pointing to a having 10

	fmt.Println(*p)
	fmt.Println(a)
}
