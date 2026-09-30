package main

import "fmt"

func change(z **int) {

	fmt.Println("function ----printing")
	fmt.Println("z pararmeter is :", z)
	fmt.Println("*z :", *z)
	fmt.Println("**z :", **z)
	fmt.Println("&z :", &z)
	fmt.Println()

	//z pointer of pointer, has adress of main pointer p
	x := 50

	*z = &x // *z now points x,not main func z

	//**z=x   --->50
	fmt.Println("_____after *z=&X_______")
	fmt.Println()
	fmt.Println("z pararmeter is :", z)
	fmt.Println("*z :", *z)
	fmt.Println("**z :", **z)
	fmt.Println("&z :", &z)

}

func main() {
	a := 10
	p := &a //*p=10

	fmt.Println()
	fmt.Println("in main----------")
	fmt.Println("value of a: ", a)
	fmt.Println("value of &a: ", &a)
	fmt.Println("value of p: ", p)
	fmt.Println("value of &p: ", &p)
	fmt.Println("value of *p: ", *p)
	fmt.Println()

	change(&p) //adress of pointer p
	fmt.Println()
	fmt.Println("______________main function _________")

	fmt.Println(*p) //now p is equal to z...so *p=**z=&x=50
	fmt.Println(a)  //a=10
}
