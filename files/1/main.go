package main

import (
	"fmt"
	"os"
)

func main() {

	myfile, _ := os.Create("student.txt")
	fmt.Println("Files created successfully")

	// defer fmt.Println("Files deleted successfully")
	// defer os.Remove("student.txt")

	defer myfile.Close()

	myfile.WriteString(`

	Name: Irfan
    Age: 20
    Country: Pakistan
	`)
	mydata, _ := os.ReadFile("student.txt")
	fmt.Println(string(mydata))

}
