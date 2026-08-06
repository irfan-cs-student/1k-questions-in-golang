package main

import (
	"fmt"
	"math"
)

// Return smallest, second smallesr number, 3rd smallest and their diference

func _3rdLargest(a [6]int) int {

	_1st, _2nd, _3rd := math.MinInt, math.MinInt, math.MinInt

	for _, value := range a {

		if value > _1st {
			_3rd = _2nd
			_2nd = _1st
			_1st = value

		} else if value < _1st && value > _2nd {

			_3rd = _2nd
			_2nd = value
		} else if value > _3rd && value < _2nd {
			_3rd = value
		}

	}

	return _3rd
}

func main() {

	nums := [6]int{2, 7, 11, 15, 3, 6}

	fmt.Println("3rd largest :", _3rdLargest(nums))

}
