package main

import "fmt"

//what the ouput ??

type Student struct {
	name string
}

func (s Student) changeName() {
	s.name = "Ali"
}
func (s *Student) altName() {
	s.name = "Ali"
}

func main() {
	s := Student{name: "Irfan"}

	s.changeName() //sends copy

	fmt.Println(s.name) //name:irfan

	s.altName()         //sends adress
	fmt.Println(s.name) //name:ali from the func alt nside

}
