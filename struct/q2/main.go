// Question 2: Employee Salary Manager
// Easy

// Create an Employee struct with the following fields:

// name → string
// age → int
// salary → int

// Tasks

// Create a slice containing 3 employees.
// Use a loop to print each employee's name and salary.
// Find the employee with the highest salary.
// Calculate the average salary of all employees.
// Increase one employee's salary by 5,000 and print the updated salary.

package main

import "fmt"

type Employee struct {
	name   string
	age    int
	salary int
}

func main() {

	// a slice containing 3 employees.

	mulazam := []Employee{

		{"irfan", 22, 7402},
		{"usman", 18, 8700},
		{"ali", 25, 2500},
	}
	fmt.Println("_______Employee name and salary__________")
	fmt.Println()

	// a loop to print each employee's name and salary.

	for i, employ := range mulazam {

		fmt.Println(i, "name: ", employ.name, " salary : ", employ.salary)
	}

	high_salary := mulazam[0]
	name := ""

	for _, employ := range mulazam {

		if employ.salary > high_salary.salary {

			high_salary.salary = employ.salary
			name = employ.name
		}
	}

	fmt.Println("name:", name, "--salary:", high_salary.salary)
}
