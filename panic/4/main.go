package main

import "fmt"

func one() {
	defer fmt.Println("1A")
	two()
	fmt.Println("1B")
}

func two() {
	defer fmt.Println("2A")
	defer func() {
		fmt.Println("2B")
		panic("PANIC-2")
	}()

	panic("PANIC-1")
}

func main() {
	defer fmt.Println("MAIN")
	one()
	fmt.Println("END")
}

//ouputa_________
/*

2b
2a
panic 2 first rewinds 2a from stack and then prints 2b becuse 2b is not in stack
panic1
a1
main

*/
