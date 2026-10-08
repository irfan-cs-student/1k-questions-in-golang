package main

import (
	"errors"
	"fmt"
)

var noMoney = errors.New("not enough money")

func buyProduct(price int, money int) (int, error) {

	if price > money || price < 0 {
		return 0, noMoney

	}

	return money - price, nil
}
func main() {

	light, err := buyProduct(100, 2)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("remaing money: ", light)

	}
}
