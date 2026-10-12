package main

import "fmt"

// Named return: defer can change the returned value.

func test() (result int) {
	result = 50

	defer func() {
		result = 100
	}()

	return
}

// Unnamed return: changing another variable doesn't change the return value.
func test2() int {
	result := 50
	defer func() { result = 100 }()
	return result
}

func main() {
	fmt.Println(test())
	fmt.Println(test2())
}
