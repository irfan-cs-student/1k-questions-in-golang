package main

import "fmt"

// Interface rejects a type

// Given:
// type Runner interface {
//     run()
// }
// Create:
// Animal
// Car
// Only Animal has run().
// Try:
// startRunning(animal)
// startRunning(car)

// Predict which line produces the compile error
// and why before running it.

type Animal struct{}
type Car struct{}

// func (a Animal) run() {
// 	fmt.Println("animals can run")
// }
func (c Car) run() {
	fmt.Println("Car can move")
}

type Runner interface {
	run()
}

func isRun(r Runner) {
	r.run()
}
func main() {
	animal := Animal{}
	car := Car{}

	car.run()
	// animal.run()

	fmt.Println("____using interfaces_______")
	fmt.Println()

	isRun(car)
	// isRun(animal)
	fmt.Print(animal)
}
