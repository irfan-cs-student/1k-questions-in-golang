package main
//open and close the files and learning  about files
import (
	"fmt"
	"os"
)

func main() {
	// 1. Create a file and write text
	file, err := os.Create("notes.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	file.WriteString("I am learning Go.")
	file.Close()

	// 2. Open the file and append text
	file, err = os.OpenFile("notes.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	file.WriteString("\nI will become a backend developer.")
	file.Close()

	// 3. Read and print the complete contents
	data, err := os.ReadFile("notes.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))

	// 4. Delete the file
	err = os.Remove("notes.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("File deleted successfully")
}
