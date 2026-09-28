// 2. Calculate Total Product Prices

package main

import "fmt"

func main() {
	products := map[string]int{
		"Laptop":   1000,
		"Mouse":    20,
		"Keyboard": 50,
		"Monitor":  200,
		"Headset":  30,
	}

	total := 0

	for _, price := range products {
		total += price
	}

	fmt.Println("total price for products:", total)

}
