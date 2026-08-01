package main

import "fmt"

//Replace negative numbers with 0

func Replace(a [8]int) (b [8]int, c [8]int) {

	for i := 0; i < len(a); i++ {

		b[i] = a[i]

		if a[i] < 0 {

			a[i] = 0

		}
		c[i] = a[i]

	}
	return b, c
}

func main() {
	nums := [8]int{12, 67, 34, -89, 51, -23, 90, 45}

	fmt.Println(Replace(nums))

}
