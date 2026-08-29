// Underlying array challenge

// Predict the output before running it, then explain why:
// a := []int{10, 20, 30, 40, 50}
// b := a[1:4]
// b[1] = 999
// fmt.Println(a)
// fmt.Println(b)

// Then modify the program so that changing b does not affect a.

package main

import "fmt"

func main() {

	a := []int{10, 20, 30, 40, 50}
	b := append([]int{}, a[1:4]...)

	b[0] = 999
	c := append([]int{}, a[1:4]...)

	fmt.Println(a) //a=[10,20,99,40]
	fmt.Println(b) //b=[20,99,40]

	fmt.Print(c)

}
