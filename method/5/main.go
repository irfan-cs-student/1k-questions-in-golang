package main

import "fmt"

// Create an Employee struct:
// Write a method increaseSalary() that increases the salary by 10%.
// Then print the employee's salary before and after calling the metho

type Employee struct {
	name   string
	salary int
}

func (emp *Employee) increaseSalary(increment int) {

	fmt.Println("___before _____", emp)
	emp.salary += emp.salary * increment / 100
}
func main() {
	employ := Employee{name: "arshad", salary: 2500}
	employ.increaseSalary(10)
	fmt.Println("___after _____", employ)

}
