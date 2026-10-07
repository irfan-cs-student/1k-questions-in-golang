package main

import "fmt"

// 2. Interface with two methods
// Create:

// type Worker interface {
//     work()
//     rest()
// }

// Create Human and Robot.
// Both must implement both methods.
// Then:
// func manage(w Worker)
// Inside manage, call both methods.
// Question: What happens if Robot implements work() but not rest()?

type Human struct{ name string }
type Robot struct{ name string }

func (h *Human) work() {
	h.name = "french"
	fmt.Println(h.name, " is working .")
}
func (h Human) rest() {
	fmt.Println(h.name, " is At rest .")

}
func (r Robot) work() {
	fmt.Println(r.name, " is working .")
}
func (r Robot) rest() {
	fmt.Println(r.name, " is At rest .")
}

type Manage interface {
	work()
	rest()
}

func isWorking(m Manage) {
	m.work()
	m.rest()
}

func main() {

	afrikan := Human{"africans"}
	agi := Robot{"AGI"}

	fmt.Println("______simple methods____")
	fmt.Println()

	afrikan.work()
	afrikan.rest()
	agi.work()
	agi.rest()
	fmt.Println()

	fmt.Println("______Using interfaces____")
	fmt.Println()

	isWorking(&afrikan)
	isWorking(agi)
	fmt.Println()

}
