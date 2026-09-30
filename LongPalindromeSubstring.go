// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func isPalindrome(s string) bool {

	for i := 0; i < len(s)/2; i++ {
		fmt.Println("s:", s, len(s))
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {

	//k:=2
	temp := ""
	pstr := ""
	if len(s) == 1 {
		return s
	}
	for k := 2; k <= len(s); k++ {
		for i := 0; i < len(s)-k+1; i++ {
			for j := i; j <= k+i; j++ {
				temp = s[i:j]
			}
			fmt.Println("temp:", temp)
			if isPalindrome(temp) {
				pstr = temp
			}
		}

	}
	return pstr
}

func main() {
	input := "abcabacdbbad"

	fmt.Println("is palindrome", longestPalindrome(input))

	fmt.Println("Hello, 世界")
}
