// Create an Employee struct with these fields:

// name → string
// age → int
// salary → int
// department → string

// Tasks:

// Create a slice containing 5 employees from different departments.
// Print all employees belonging to the "IT" department.
// Calculate the total salary of employees in the IT department.
// Find the highest-paid employee in the entire college.
// Increase the salary of every IT employee by 10%.

package main

import "fmt"

type Employee struct {
	name       string
	age        int
	salary     int
	department string
}

func main() {

	professors := []Employee{

		{"ali", 40, 4000, "it"},
		{"usman", 80, 455, "cs"},
		{"shakeel", 45, 858, "ai"},
		{"falak", 56, 757, "it"},
		{"hyder", 58, 957, "it"},
		{"rizwan", 80, 740, "se"},
	}

	// Print all employees belonging to the "IT" department.
	fmt.Println("IT depart employee________")
	fmt.Println()

	for _, emp := range professors {

		if emp.department == "it" {

			fmt.Println("Employee name: ", emp.name, " ---salary:  ", emp.salary)
		}
	}

	// Calculate the total salary of employees in the IT department.
	total := 0
	for _, prof := range professors {
		if prof.department == "it" {

			total += prof.salary
		}
	}
	fmt.Println("it department total salary:", total)
	fmt.Println()

	// Find the highest-paid employee in the entire college.
	high_pay_employ := professors[0]

	for _, emp := range professors {
		if emp.salary > high_pay_employ.salary {
			high_pay_employ = emp
		}
	}
	fmt.Println("high pay employ: ", high_pay_employ.name, "--salary:", high_pay_employ.salary)
	fmt.Println()

	// Increase the salary of every IT employee by 10%.

	for i := range professors {
		if professors[i].department == "it" {

			professors[i].salary = int(float64(professors[i].salary) * 1.10)
		}
	}
	fmt.Println("it profsoors salay after 10% increments")

	for i := range professors {
		if professors[i].department == "it" {
			fmt.Print("name: ", professors[i].name)
			fmt.Println("--salary: ", professors[i].salary)
		}
	}
	fmt.Println()

}
