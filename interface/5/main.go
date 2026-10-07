package main

import "fmt"

// Create:
// type Speaker interface {
//     speak()
// }
// Create:
// Human,Dog,Robot
// All implement speak() differently.
// Then create:
// speakers := []Speaker{
//     Human{},
//     Dog{},
//     Robot{},
// }
// Loop through the slice and call:
// s.speak()
// This one is very important for understanding why interfaces are useful.

type Human struct{}
type Dog struct{}
type Robot struct{}

func (h Human) speak() {
	fmt.Println("he can speak")
}
func (d Dog) speak() {
	fmt.Println("dog can speak")
}
func (r Robot) speak() {
	fmt.Println("robot can talk")
}

type Speaker interface {
	speak()
}

func canSpeak(s Speaker) {
	s.speak()
}

func main() {
	human := Human{}
	dog := Dog{}
	robot := Robot{}

	canSpeak(human)
	canSpeak(dog)
	canSpeak(robot)

	talk := []Speaker{
		Human{}, Dog{}, Robot{},
	}

	for _, s := range talk {
		s.speak()
	}

}
