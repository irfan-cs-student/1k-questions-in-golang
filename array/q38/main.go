package main

import "fmt"

// Rotate array to the left by 'steps'
func moveAhead(a [6]int, steps int) (result [6]int) {

	steps = steps % len(a)
	// Handle steps greater than array length

	for i := range a {
		result[i] = a[(i+steps)%len(a)]
	}

	return
}

func main() {
	num := [6]int{1, 2, 3, 4, 5, 6}

	var steps int
	fmt.Print("Enter steps: ")
	fmt.Scan(&steps)

	fmt.Println("Original:", num)
	fmt.Println("Rotated :", moveAhead(num, steps))
}
