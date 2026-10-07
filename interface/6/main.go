package main

import "fmt"

// Interface hides other methods
type Runner interface {
	run()
}

// Animal has three methods
type Animal struct {
	name string
}

func (a Animal) run() {
	fmt.Println(a.name, "is running")
}

func (a Animal) eat() {
	fmt.Println(a.name, "is eating")
}

func (a Animal) sleep() {
	fmt.Println(a.name, "is sleeping")
}

// Runner only gives access to run()
func test(r Runner) {
	r.run()

	// r.eat()   // rror: Runner has no method eat()
	// r.sleep() // Error: Runner has no method sleep()
}

func main() {
	animal := Animal{name: "Dog"}

	test(animal)

	// Direct Animal value can access all methods
	animal.run()
	animal.eat()
	animal.sleep()
}
