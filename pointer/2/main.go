package main

import "fmt"

func change(a *int, b *int) {
	*a = *a + 5 //a--10+5=15

	temp := a //temp=adrees of a,not value of a
	a = b     //a poits y becuse storing y adress geting from b
	b = temp  //b points a ,

	//here
	//a nd b has no int values of x ,y but long adreess valuesa
	//*a and *b has actual variable values not index ,but these points to x,y in main func

	*b = *b + 10 //b=x value---15+10=25 pointing to x becuase of chainging adreese from temp
	*a = *a * 2  //a=y value---20*2=400 poiting to y becuse a now points y
}

func main() {
	x := 10
	y := 20

	p := &x
	q := &y

	change(p, q)

	fmt.Println(x)  //25
	fmt.Println(y)  //40
	fmt.Println(*p) //25
	fmt.Println(*q) //40
}
