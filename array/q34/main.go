package main

import (
	"fmt"
	"math"
)

// Return smallest, second smallesr number, 3rd smallest and their diference

func extremeNum(a [6]int) (smallest, _2ndSmall, _3rdSmall int) {

	smallest, _2ndSmall, _3rdSmall = math.MaxInt, math.MaxInt, math.MaxInt

	for _, value := range a {

		if value < smallest {
			_3rdSmall = _2ndSmall
			_2ndSmall = smallest
			smallest = value

		} else if value > smallest && value < _2ndSmall {

			_3rdSmall = _2ndSmall
			_2ndSmall = value
		} else if value < _3rdSmall && value > _2ndSmall {
			_3rdSmall = value
		}

	}

	return
}

func main() {

	nums := [6]int{2, 1, -7, 0, 4, -2}

	smallest, _2ndSmall, _3rdSmall := extremeNum(nums)
	fmt.Println("smallest:", smallest)
	fmt.Println("2ndsmall:", _2ndSmall)
	fmt.Println("3rdsmall:", _3rdSmall)

}
