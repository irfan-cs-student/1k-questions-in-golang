package main

import "fmt"

type Product struct {
	name     string
	price    int
	quantity int
}

type Store []Product

func (s Store) discountAll(persent int) {

	for i := range s {

		s[i].price -= s[i].price * persent / 100

	}
}
func main() {
	store := Store{
		{"Laptop", 100000, 2},
		{"Mouse", 2000, 0},
		{"Keyboard", 5000, 3},
		{"USB", 120, 5},
	}
	copy := store
	fmt.Println(copy)

	store.discountAll(20)
	fmt.Println(store)
	fmt.Println(copy)
}
