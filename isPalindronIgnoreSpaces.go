// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(str string) bool {
	teststr := []rune(str)
	l := len(str)
	for i := 0; i < l/2; i++ {
		if teststr[i] == teststr[l-(i+1)] {
			continue
		} else {
			return false
		}
	}
	return true
}

func isCleaned(str string) string {
	cleaned := ""
	for _, ch := range strings.ToLower(str) {
		if unicode.IsLetter(ch) {
			cleaned = cleaned + string(ch)
		}
	}
	fmt.Println("\nTest:", string(cleaned))
	return string(cleaned)
}

func main() {
	str := "A man, a plan , a canal :Panama"

	fmt.Println("isPalindrome", isPalindrome(isCleaned(str)))
	fmt.Println("Hello, 世界")
}
