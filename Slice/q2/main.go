package main

import "fmt"

// Create a slice with make. Print its length and capacity. Append 3 elements
// and print length/capacity again. Explain through comments why they changed.

func main() {
	a := make([]int, 5)
	fmt.Println(a)
	fmt.Println("capacity:", cap(a))
	fmt.Println("lenght:", len(a))

	//expect:[0,0,0,0,0],capacity 5,lenght 5 same

	//appending the elements
	a = append(a, 3, 4, 5)
	fmt.Println(a)
	fmt.Println("capacity:", cap(a))
	fmt.Println("lenght:", len(a))

	//expect:[0,0,0,0,0,3,4,5],lenght 8,capacity 10
	// because if lenght reaches capacity ,capacity increase automatciaally
	// with rule 2times of prevous lenght : 2*5=10

}
