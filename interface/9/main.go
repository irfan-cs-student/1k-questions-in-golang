package main

import "fmt"

//struct values changing

type Animal struct{ name string }

func (a Animal) eat() {
	fmt.Println(a.name, "can eat grass")
}

type eating interface {
	eat()
}

func eaten(e eating) {
	e.eat()
}
func main() {

	//goat := Animal{} we will not declare as

	// others ways

	(Animal{"goat"}).eat()
	eaten(Animal{"sheep"})
}
