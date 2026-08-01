package main

import "fmt"

//Sum of elements at even and odd index

func evenOdd(a [8]int) (int, int) {

	evenSum, oddSum := 0, 0

	for i := 0; i < len(a); i++ {

		if i%2 == 0 {

			evenSum += a[i]
		} else {
			oddSum += a[i]

		}

	}
	return evenSum, oddSum
}

func main() {

	nums := [8]int{5, 2, 7, 2, 9, 2, 4, 8}
	evenSum, oddSum := evenOdd(nums)

	fmt.Println("sum of elements at even index :", evenSum)
	fmt.Println("sum of elements at odd index:", oddSum)

}
