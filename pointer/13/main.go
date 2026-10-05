package main

import "fmt"

//pointers in structs
type students struct {
	name string
	age  int
}

func main() {
	//simple struk value
	stud_1 := students{"irfan", 22}
	p := stud_1

	//slice struct value
	class := []students{
		{"irfan", 22},
		{"", 33},
	}
	a := &class

	fmt.Println(p.name)
	fmt.Println(p.age)

	fmt.Println("slice struct________")
	fmt.Println((*a)[0].name)
	fmt.Println((*a)[0].age)

}
