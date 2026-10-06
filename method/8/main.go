package main

import "fmt"

type Temperature float64

func (t Temperature) toFahrenheit() float64 {
	return (float64(t) * 9 / 5) + 32
}

func (t Temperature) toCelsius() float64 {
	return (float64(t) - 32) * 5 / 9
}

func main() {
	c := Temperature(25)

	fmt.Println(c.toFahrenheit())
	fmt.Println(c.toCelsius())
}
