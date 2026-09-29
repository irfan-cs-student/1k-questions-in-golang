package main

import "fmt"

func change(a *int, b *int) {

	*a = *a + 10 //a--5+10=15
	*b = *a + *b //b---15+15=30
	*a = *b - *a //a---30-30=0
}

func main() {
	x := 5

	p := &x
	q := &x

	change(p, q)

	fmt.Println(x)  //=
	fmt.Println(*p) //*p=x=0
	fmt.Println(*q) //*q=x=0
}
