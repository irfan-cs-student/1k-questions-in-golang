package main

import "fmt"

type Counter struct {
	value int
}

func (c *Counter) increment() {
	fmt.Println(c.value)

	c.value++ //6 for 1st calling,7 for second calling
	fmt.Println(c.value)

}

func main() {
	c := Counter{value: 5}

	c.increment() //6
	c.increment() //7

	fmt.Println(c.value)
}
