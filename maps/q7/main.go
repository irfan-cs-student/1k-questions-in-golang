// 4. Count Word Frequency
package main

import "fmt"

func main() {
	words := []string{"go", "python", "go", "java", "python", "go"}

	count := make(map[string]int)

	for _, words := range words {
		count[words] = count[words] + 1

		//upper expressions equal to __count[words]++
	}
	fmt.Print(count)

}
