// Predict the output first:

// a := make([]int, 2, 5)
// a[0] = 10
// a[1] = 20

// b := a[:1]

// b = append(b, 99)

// fmt.Println(a)
// fmt.Println(b)

// Then change the capacities/append operations and
// investigate when append modifies the same underlying array
//  and when it causes allocation of a new one.

package main

import "fmt"

func main() {

	a := make([]int, 2, 5)
	a[0] = 10
	a[1] = 20

	//a=[10,20,0]
	b := a[:1]

	//b=[10]
	b = append(b, 99)

	//b=[10,99]
	fmt.Println(a) //[10,20]>> [10,99]
	fmt.Println(b) //[10,99]

	b = append(b, 2, 3)
	fmt.Println(a)
	fmt.Println(b)

	b = append(b, 4, 5, 6, 7, 8)
	fmt.Println(a)
	fmt.Println(b)

}
