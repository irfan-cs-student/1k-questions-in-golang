package main

import "fmt"

// Merge two already-sorted slices
func merge(a []int, b []int) []int {

	var c []int

	i := 0
	j := 0

	// Compare elements from both slices
	for i < len(a) && j < len(b) {

		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}

	// Add remaining elements from a
	for i < len(a) {
		c = append(c, a[i])
		i++
	}

	// Add remaining elements from b
	for j < len(b) {
		c = append(c, b[j])
		j++
	}

	return c
}

func main() {

	a := []int{1, 4, 7, 10}
	b := []int{2, 3, 8, 9}

	result := merge(a, b)

	fmt.Println("Merged slice:", result)
}
