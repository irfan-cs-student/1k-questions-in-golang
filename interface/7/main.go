package main

import "fmt"

type Movable interface {
	move()
}

type Bird struct{}
type Car struct{}
type Human struct{}

func (b Bird) move() {
	fmt.Println("Bird is flying")
}

func (c Car) move() {
	fmt.Println("Car is driving")
}

func (h Human) move() {
	fmt.Println("Human is walking")
}

func moveSomething(m Movable) {
	m.move()
}

func main() {
	moveSomething(Bird{})
	moveSomething(Car{})
	moveSomething(Human{})
}
