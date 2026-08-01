package main

import "fmt"

//seperate even and odd

func evenOdd(a [8]int) (even [9]int, odd [9]int) {

	e1, e2 := 0, 0

	for i := 0; i < len(a); i++ {

		if a[i]%2 == 0 {

			even[e1] = a[i]
			e1++

		} else {

			odd[e2] = a[i]
			e2++
		}

	}
	return even, odd
}

func main() {

	nums := [8]int{5, 2, 7, 2, 9, 2, 4, 8}
	even, odd := evenOdd(nums)

	fmt.Println("even:", even)
	fmt.Println("odd:", odd)

}
