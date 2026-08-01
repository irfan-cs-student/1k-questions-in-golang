package main

import "fmt"

//reverse array as first half becomes second half and vice versa
//means : replced first half with last half

func Replace(a [8]int) (b [8]int) {

	//seond half into first half new  array
	index := 0
	for i := len(a) / 2; i < len(a); i++ {

		b[index] = a[i]
		index++
	}

	for i := 0; i < len(a)/2; i++ {

		b[index] = a[i]
		index++
	}

	return b
}

func main() {
	nums := [8]int{12, 67, 34, -89, 51, -23, 90, 45}

	fmt.Println("replced first half with last half", Replace(nums))

}
