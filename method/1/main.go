package main

import "fmt"

// Create a Rectangle struct with width and height.
// Write a method area() that returns the rectangle's area.
type Rectangle struct {
	height, width int
}

func (shape Rectangle) area() int {

	area := shape.height * shape.width
	return area

}

func main() {

	shape := Rectangle{4, 8}
	fmt.Println("area:", shape.area())
}
