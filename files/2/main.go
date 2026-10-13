// Question 2 — Append and Delete a File
package main

import (
	"fmt"
	"os"
)

func main() {

	file, error := os.Create("notes.txt")
	if error != nil {
		fmt.Println(error)
	}

	file.WriteString(`
I am learning Go.
I will become a backend developer.`)

	data, _ := os.ReadFile("notes.txt")
	fmt.Println(string(data))

	file.Close()
	fmt.Println("File deleted successfully")

	// os.Remove("notes.txt")
}
