// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func removeDuplicates(rawText string) string {
	words := strings.Fields(rawText)
	seen := make(map[string]bool)
	result := []string{}

	for _, word := range words {
		_, exists := seen[word]
		if !exists {
			seen[word] = true
			result = append(result, word)
		}
	}
	return strings.Join(result, " ")
}

func main() {
	text := " this is is a test test paragraph"

	cleaned := removeDuplicates(text)

	fmt.Println("Cleaned Text: ", cleaned)

	fmt.Println("Hello, 世界")
}
