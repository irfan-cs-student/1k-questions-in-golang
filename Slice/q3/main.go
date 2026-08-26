// Slicing a slice
// Given:
// nums := []int{10, 20, 30, 40, 50, 60, 70, 80}
// Create a new slice containing 30, 40, 50, 60 without manually writing those values.

package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40, 50, 60, 70, 80}

	b := nums[2:6]
	fmt.Println(b)

	fmt.Println("lenght of nums:", len(nums), " ", "lenght of b", len(b))
	fmt.Println("capacity of nums:", cap(nums), " ", "capacity of b", cap(b))

	//output: from 2 index to before 6,but not 6 itself

	b[0] = -1
	fmt.Println(b)

	//output: -1,40,50

	b = append(b, -10)
	fmt.Println(b)
	fmt.Println("lenght of nums:", len(nums), " ", "lenght of b", len(b))
	fmt.Println("capacity of nums:", cap(nums), " ", "capacity of b", cap(b))

	//output:-1,40,50,-10
	//index 6 of nums now -10

	fmt.Println(nums)

	//now appending more with b
	b = append(b, -11, -12, -13, -14, 15, -16)
	fmt.Println(b)
	fmt.Println(nums)
	fmt.Println("lenght of nums:", len(nums), " ", "lenght of b", len(b))
	fmt.Println("capacity of nums:", cap(nums), " ", "capacity of b", cap(b))

}
