package main

import "fmt"

func change(z **int) {

	//z pointer of pointer, has adress of main pointer p
	x := 50

	*z = &x // *z now points x,not main func z

	//**z=x   --->50
}

func main() {
	a := 10
	p := &a //*p=10

	change(&p) //adress of pointer p

	fmt.Println(*p) //now p is equal to z...so *p=**z=&x=50
	fmt.Println(a)  //a=10
}
