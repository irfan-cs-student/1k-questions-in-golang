package main

//pointers in array . pointers cahnge oringanl data
//simple value of index change doesn't change the orignal value
//without pointer

import "fmt"

func alter(arr [5]int) {

	if arr[0] == -1 {
		for i := 0; i < len(arr); i++ {
			(arr)[i] += 2
		}
		fmt.Println("2nd function printing________")
		fmt.Println("updated array :", arr)
		return
	}
	arr[0] = -1

	fmt.Println("function printing ___")
	fmt.Println("array:", arr)

}
func main() {
	s := [5]int{3, 4, 5, 3, 2}
	p := &s

	fmt.Println(s[1])
	fmt.Println((*p)[1])

	alter(s)
	alter(*p)
	(*p)[0] = -1
	alter(s)
	alter(*p)

	fmt.Println()
	fmt.Println("main array:", s)
	fmt.Println("main array:", *p)

}
