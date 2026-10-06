package main

import "fmt"

//  Create a Student struct with name and age.
//  Write a method introduce() that prints:
// My name is Irfan and I am 22 years old.

type Student struct {
	name string
	age  int
}

func (s Student) introduce() {
	fmt.Println("my name:", s.name)
	fmt.Println("my age:", s.age)
}

func main() {

	students := Student{"irfan", 22}
	students.introduce()

}
