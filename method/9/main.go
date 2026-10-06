package main

import "fmt"

type Product struct {
	name  string
	price int
}

func (p *Product) discount(percent int) {

	p.price -= p.price * percent / 100

}

func main() {

	products := []Product{
		{"Laptop", 100000},
		{"Mouse", 2000},
		{"Keyboard", 5000},
	}

	for i := range products {
		products[i].discount(20)
	}
	fmt.Println(products)

}
