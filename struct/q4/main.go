// Create an Order struct with these fields:

// name     → string
// product  → string
// price    → int
// quantity → int

// Tasks
// Create a slice containing 5 orders.
// Print each customer's name, product, and total cost.
// Total cost = price × quantity
// Find the order with the highest total cost.
// Calculate the total revenue from all orders.
// Increase the price of "Laptop" products by 10%.
// Print the updated laptop prices.

package main

import "fmt"

type Order struct {
	name     string
	product  string
	price    int
	quantity int
}

func main() {
	orders := []Order{
		{"Irfan", "Laptop", 80000, 2},
		{"Ali", "Mouse", 2500, 3},
		{"Usman", "Keyboard", 5000, 2},
		{"Ahmed", "Laptop", 90000, 1},
		{"Sara", "Monitor", 30000, 2},
	}
	// Print each customer's name, product, and total cost.
	fmt.Println("Gahak/customers data-------")
	fmt.Println()

	high_price_product := orders[0]
	revenue := 0

	for _, gahak := range orders {
		fmt.Println("customers name:", gahak.name)
		fmt.Println("product:", gahak.product)
		fmt.Println("quantity:", gahak.quantity)
		fmt.Println("price pr 1 item :", gahak.price)
		fmt.Println("total price : ", gahak.quantity*gahak.price)
		fmt.Println()

		{
			// Find the order with the highest total cost.
			if gahak.price > high_price_product.price*high_price_product.quantity {
				high_price_product = gahak
			}
		}
		{ // Calculate the total revenue from all orders.

			revenue += gahak.price * gahak.quantity
		}
	}
	fmt.Println("total revenue:", revenue)

	fmt.Println("costly item/product in all: ", high_price_product.product,
		" --price:", high_price_product.price)

	fmt.Println()

	for i := range orders {
		if orders[i].product == "Laptop" {
			orders[i].price += int(float64(orders[i].price) * 1.10)
		}
	}
	for _, prod := range orders {
		if prod.product == "Laptop" {
			fmt.Println("laptop updated price:", prod.price)

		}
	}

}
