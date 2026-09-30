package main

import "fmt"

func modify(p **int) {

	//here
	//p has adress of main p variable
	//*p value of (=*P main variable)
	//**p pointing to a, **P=&p=a=10

	**p = 100
	//now **p=a=100
}

func main() {
	a := 10
	p := &a //*p=10

	modify(&p) //gets **p=&p --->giving/pointing  adress of a=10,changes to 100

	fmt.Println(a)  //100
	fmt.Println(*p) //100 becuase *p =value of (a) variable
}
