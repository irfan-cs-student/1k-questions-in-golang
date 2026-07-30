package main

import "fmt"

//Find a specific number

func main() {

	nums := [...]int{67, 77, -30, 40, 50, -47, 88, -99, 21}
	var specific_num_index int
	found := false

	for index, value := range nums {

		if 401 == value {

			specific_num_index = index
			found = true
		}
	}

	if found {
		fmt.Println("Specific number fount at index :", specific_num_index)
	} else {
		fmt.Println("number not found")

	}
}
