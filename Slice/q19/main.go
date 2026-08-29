// nums := []int{7, 3, 9, 3, 1, 7, 5}
// reverse and sort it

package main

import "fmt"

func revSort(a []int) {

	left, right := 0, len(a)-1
	for left < right {
		a[left], a[right] = a[right], a[left]

		left++
		right--
	}
	fmt.Println("reverse slice:", a)

	//sorting

	for i, _ := range a {

		for index := 0; index < len(a)-1-i; index++ {

			if a[index] > a[index+1] {

				a[index], a[index+1] = a[index+1], a[index]
			}
		}

	}
	fmt.Println("sorted reverse slics:", a)

	//removing duplicates

	var b []int
	for _, value := range a {
		found := false
		for _, element := range b {
			if value == element {
				found = true
				break
			}
		}
		if !found {

			b = append(b, value)
		}
	}
	fmt.Println("remove duplicates:", b)
}
func main() {
	nums := []int{7, 3, 9, 3, 1, 7, 5}
	fmt.Println("orignal slice:", nums)

	revSort(nums)

}
