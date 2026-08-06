package main

import "fmt"

// Rotate array one position to the left
func moveElement(a [6]int) (result [6]int) {

	temp := a[0]

	for index := range a {

		if index == len(a)-1 {
			result[index] = temp
		} else {
			result[index] = a[index+1]
		}
	}
	return
}

// Rotate array to the left by any number of steps
func moveAhead(a [6]int, steps int) (result [6]int) {

	steps = steps % len(a) // Handle steps larger than array size

	tempRange := steps

	for index := range a {

		if tempRange < len(a) {
			result[index] = a[tempRange]
		} else {
			result[index] = a[tempRange-len(a)]
		}

		tempRange++
	}

	return
}

func main() {

	num := [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println("Original array:", num)
	fmt.Println("Move 1 step:", moveElement(num))

	var steps int
	fmt.Print("Enter steps: ")
	fmt.Scan(&steps)

	fmt.Println("Original array:", num)
	fmt.Println("Move", steps, "steps:", moveAhead(num, steps))
}
