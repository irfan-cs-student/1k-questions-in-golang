package main

import "fmt"

// Two types, one interface

// Create:

// Animal
// Robot
// Both must satisfy:
// type Runner interface {
//     run(speed int)
// }
// Create:
// func startRunning(r Runner, speed int)
// Call it with both types.
// Goal: Understand how different types satisfy the same interface.

type Animal struct{}
type Robot struct{}

func (a Animal) run(name string, speed int) {
	fmt.Println(name, "runs at speed of: ", speed)
}
func (a Robot) run(name string, speed int) {
	fmt.Println(name, "runs at speed of: ", speed)
}

type Runner interface {
	run(name string, speed int)
}

func startRunning(r Runner, name string, speed int) {
	r.run(name, speed)
}

func main() {

	cat := Animal{}
	car := Robot{}

	startRunning(cat, "cat", 12)
	startRunning(car, "car", 150)

}
