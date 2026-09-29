// Employee Manager with Functions
package main

import "fmt"

type Employee struct {
	name       string
	age        int
	salary     int
	department string
}

func printEmployees(employee []Employee) {
	for _, emp := range employee {
		fmt.Println("name:", emp.name, "--salary:", emp.salary)

	}
}
func highestSalary(employees []Employee) Employee {

	highest := employees[0]
	for _, emp := range employees {
		if emp.salary > highest.salary {

			highest = emp
		}
	}
	return highest
}
func totalSalary(employee []Employee) int {

	totalSalary := 0
	for _, salary := range employee {

		totalSalary += salary.salary
	}
	return totalSalary
}
func increaseSalary(employee []Employee, department string) {

	for i := range employee {
		if employee[i].department == department {
			employee[i].salary += int(float64(employee[i].salary) * 1.01)
		}
	}

}
func main() {

	employees := []Employee{
		{"Irfan", 22, 74000, "IT"},
		{"Usman", 28, 65000, "HR"},
		{"Ali", 25, 82000, "IT"},
		{"Ahmed", 30, 58000, "Finance"},
		{"Sara", 27, 91000, "IT"},
	}
	fmt.Println()
	printEmployees(employees)
	fmt.Println()

	highest := highestSalary(employees)
	fmt.Println("name:", highest.name)
	fmt.Println("hightest_salary:", highest.salary)

	total := totalSalary(employees)
	fmt.Println("total_salary:", total)
	fmt.Println()

	// return changed/increase salarties for department

	increaseSalary(employees, "it")
	fmt.Println("employees from it after incresing salary by 10%")

	for _, emp := range employees {

		if emp.department == "IT" {
			fmt.Print("name:", emp.name)
			fmt.Println("  salary:", emp.salary)

		}
	}

}
