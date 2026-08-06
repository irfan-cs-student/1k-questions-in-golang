package main

import "fmt"

// MostFrequent number
func frequentNum(a [5]int) (mostFreq, preCount int) {

	mostFreq, preCount = a[0], 0

	for _, value := range a {

		count := 0

		for _, num := range a {
			if value == num {

				count++
			}

		}
		if count > preCount {

			mostFreq = value
			preCount = count
		}

	}
	return

}

func main() {

	nums := [5]int{1, 2, 2, 9, 2}
	mostFreq, count := frequentNum(nums)

	fmt.Print("MostFrequent Number : ", mostFreq, "\n repeated count:", count)

}
