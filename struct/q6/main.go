//  Student Result Manager

package main

import "fmt"

type Student struct {
	name  string
	age   int
	marks int
	grade string
}

func printStudent(pupils []Student) {

	for _, student := range pupils {
		fmt.Println("name: ", student.name, "  --marks:", student.grade)

	}
}
func highestMarks(pupils []Student) Student {

	highestMarks := pupils[0]
	for _, student := range pupils {
		if highestMarks.marks < student.marks {
			highestMarks = student
		}
	}
	return highestMarks
}
func averageMarks(s []Student) float64 {
	total := 0
	for _, Student := range s {
		total += Student.marks
	}
	averageMarks := total / (len(s))
	return float64(averageMarks)
}
func passedStudents(s []Student, limit int) {

	for _, Student := range s {
		if Student.marks >= limit {
			fmt.Println(Student.name, "--marks:", Student.marks)
		}
	}

}
func alterMarks(s []Student, nameOfstudent string) {
	for i := range s {
		if s[i].name == nameOfstudent {
			fmt.Println(s[i].name, "  --marks:=", s[i].marks)

			s[i].marks += 5
			fmt.Println("after modifying the marks")
			fmt.Println(s[i].name, "  --marks:=", s[i].marks)
		}
	}
}
func main() {
	students := []Student{
		{"Irfan", 22, 85, "A"},
		{"Ali", 20, 72, "B"},
		{"Usman", 23, 91, "A"},
		{"Ahmed", 21, 48, "F"},
		{"Sara", 22, 67, "C"},
	}
	//print students with name
	fmt.Println()
	printStudent(students)
	fmt.Println()

	// Return the student with the highest marks.
	topper := highestMarks(students)
	fmt.Println(topper.name, "  marks:", topper.marks, "--grad:", topper.grade)

	//print average
	fmt.Println("average marks:", averageMarks(students))

	//Print students whose marks are 50 or above.
	fmt.Println("Print students whose marks are 60 or  above")

	passedStudents(students, 60)

	// Increase the marks of one student by 5.
	fmt.Println(" Increase the marks of one student by 5")

	alterMarks(students, "Ahmed")

}
