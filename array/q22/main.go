package main

import "fmt"

//multiply  with 2
// and also make a new array where element multiply with 3 where values are 50-100 in between

func Replace(a [8]int) (b [8]int, c [8]int) {

	for i := 0; i < len(a); i++ {

		c[i] = a[i]
		if a[i] >= 50 && a[i] <= 100 {
			c[i] = c[i] * 3
		}
		b[i] = a[i] * 2

	}
	return b, c
}

func main() {
	nums := [8]int{12, 67, 34, -89, 51, -23, 90, 45}

	fmt.Println(Replace(nums))

}
