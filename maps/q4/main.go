//

package main

import "fmt"

func main() {

	scores := map[string]int{
		"Ali":   90,
		"Ahmed": 80,
		"Sara":  95,
	}

	//loops through map
	for name, value := range scores {

		if value > 90 {
			fmt.Println(name, " = ", value)
		}
	}
}
