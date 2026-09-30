package main

import "fmt"

func test(p **int) {

	//**p=value of a
	//*p=adress of a
	x := 30

	**p = 20 //**P=a=20
	*p = &x
	//*p=stores adress of a i.e; p---->a
	// now store adress changes to x
	//p--->x
	**p = 40
	//value of x=**P=40
}

func main() {
	a := 10
	p := &a //*p=10

	test(&p) //pasing adrees of p pointer(pointing to a i.e; p----->a)

	fmt.Println(a) //a=20
	// before changing stores adress a got value of 20

	fmt.Println(*p) //40
	//*p is now value of x
	//*p=x varialbe value  now its 40

	//how x survive as its not global its func local variable
	//x survives because its address escapes the function through p,
	// so Go's compiler moves/keeps x on the heap. that( i asked from gpt for concept )

	// remeber!!!

	// p=adrees of a,(&a),
	// *p=value of a,**z=value of a,
	// z=&p,*z=adress of orignal variable,
	// **z=value of a,&z=adress of z,
	// z=adress of pointer

}
