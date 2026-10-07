package main

import "fmt"

// 4. Different implementations
// Create:
// type Runner interface {
//     run()
// }
// Animal.run() should print:
// Animal is running
// Robot.run() should print:
// Robot is running mechanically
// Then:
// func start(r Runner)
// Call it with both.
// Important: Same interface method, completely different implementation.

type Animal struct{ name string }
type Robot struct{ name string }

func (a Animal) run() {

	fmt.Println(a.name, " can run")
}
func (r Robot) run() {

	fmt.Println(r.name, " can move")
}

type Runner interface {
	run()
}

func startRunning(r Runner) {
	r.run()
}

func main() {

	cat := Animal{"cat"}
	agi := Robot{"AGI"}

	cat.run()
	agi.run()

	fmt.Println("____printing using interfacs______")
	startRunning(cat)
	startRunning(agi)

}
