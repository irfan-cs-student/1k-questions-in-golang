package main

import "fmt"

//returning interfaces

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
func getAnimal() Animal {

	return Animal{"goat"}
}
func main() {

	goat := getAnimal()

	goat.eat()
	eaten(goat)
}
