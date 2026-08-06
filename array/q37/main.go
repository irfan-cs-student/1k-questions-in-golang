package main

import "fmt"

//Rotate array one position to the left

func moveElement(a [6]int) (result [6]int) {

	temp := a[0]

	for index, _ := range a {

		if index == len(a)-1 {
			result[index] = temp

		} else {
			result[index] = a[index+1]

		}

	}
	return

}
func main() {
	num := [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println("Original array:", num)
	fmt.Println("Updated array:", moveElement(num))

}
