package main

import "fmt"

func modify(a **int, b int) {
	fmt.Println("value of a:", a)
	fmt.Println("value of &a:", &a)
	fmt.Println("value of *a:", *a)
	fmt.Println("value of b:", b)

	**a = 7
	b = 6
	fmt.Println("value of a:", a)
	fmt.Println("value of b:", b)

}

func main() {

	a := 10
	b := 20
	p := &a

	modify(&p, b)

	fmt.Println("value of p:", p)
	fmt.Println("value of &p:", &p)
	fmt.Println("value of *p:", *p)

	fmt.Println("value of a:", a)
	fmt.Println("value of b:", b)

}
