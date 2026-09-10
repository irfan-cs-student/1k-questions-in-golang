// 2. Word Frequency Analyzer

// Given:

// text := "go is fast and go is simple and go is powerful"

// Create a program that:

// Splits the text into words.
// Stores each word's frequency in a map.
// Prints the frequency of every word.
// Finds the most frequently occurring word

package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "go is fast and go is simple and go is powerful"

	// Convert sentence into words
	words := strings.Fields(text)

	// Store word and its count
	count := make(map[string]int)

	// Count each word
	for _, word := range words {
		count[word]++
	}

	// Print every word and its count
	for word, number := range count {
		fmt.Println(word, ":", number)
	}

	// Find the most common word
	topWord := ""
	topCount := 0

	for word, number := range count {

		if number > topCount {
			topCount = number
			topWord = word
		}
	}

	fmt.Println("Most common word:", topWord)
	fmt.Println("Count:", topCount)
}
