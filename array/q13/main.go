package main

import "fmt"

//Find duplicate values

func duplicate(array [9]int) {

	for a := 0; a <= len(array)-1; a++ {

		for b := a + 1; b <= len(array)-1; b++ {

			if array[b] == array[a] {
				fmt.Print(array[a], " ")
			}
		}
	}

}
func dupElement(a [9]int, value int, size int) bool {

	for i := 0; i < size; i++ {

		if a[i] == value {
			return true
		}

	}
	return false
}
func findduplicate(array [9]int) [9]int {

	index := 0
	var result [9]int
	for a := 0; a <= len(array)-1; a++ {

		for b := a + 1; b <= len(array)-1; b++ {

			if array[b] == array[a] {

				value := array[b]
				isRepeate := dupElement(result, value, index)

				if isRepeate == false {
					result[index] = array[b]
					index++

				}

			}
		}
	}
	return result

}

func main() {

	nums := [9]int{10, 20, 30, 40, 50, 10, 20, 20, 50}
	fmt.Print("Duplicate values: ")
	duplicate(nums)
	fmt.Println()
	fmt.Print("Duplicaten: ", findduplicate(nums))

}
