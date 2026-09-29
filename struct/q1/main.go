// Create a Student struct with these fields:

// name → string

// age → int

// marks → int

// Tasks:

// Create a slice containing 3 students.

// Print the name and marks of each student using a loop.

// Find the student with the highest marks.

// Calculate the average marks of all students.

// Change the marks of one student and print the updated value

package main

import "fmt"

type student struct {
	name  string
	age   int
	marks int
}

func main() {

	//using struct as slice data type
	students := []student{
		{name: "irfan", age: 22, marks: 80},
		{name: "Ahmed", age: 21, marks: 90},
		{name: "Irfan", age: 22, marks: 95},
	}

	//printing all pupils
	for i, pupil := range students {

		fmt.Println("pupil:", i)
		fmt.Println("name :", pupil.name)
		fmt.Println("age :", pupil.age)
		fmt.Println("marks :", pupil.marks)
		fmt.Println()
	}

	//finding pupils with heighest marks
	heightest := students[0]

	for _, pupils := range students {
		if pupils.marks > heightest.marks {
			heightest = pupils
		}
	}
	fmt.Println()
	fmt.Println("name :", heightest.name, " marks", heightest.marks)

	//calculate of all students summ all marks and devide total pupils

	total := 0
	for _, student := range students {
		total += student.marks
	}
	average := float64(total) / float64(len(students))
	fmt.Println("average: ", average)
	fmt.Println()

	//update value
	students[0].marks = 12
	fmt.Println("name:", students[0].name)
	fmt.Println("marks:", students[0].marks)

}
