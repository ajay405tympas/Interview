// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func validateSequence(dataSeq []string) []int32 {
	var balance int = 0

	result := make([]int32, len(dataSeq))

	fmt.Println("dataSeq:", dataSeq)

	for i := range dataSeq {
		charStr := dataSeq[i]
		fmt.Println("charStr:", charStr)

		for j := range charStr {
			char := charStr[j]
			fmt.Println("char:", string(char))
			if char == '(' {
				balance++
			} else {
				balance--
			}
		}

		if balance == 0 {
			result[i] = 1
		} else {
			result[i] = 0
		}
		balance = 0
	}

	return result
}

func main() {
	var str []string
	var n int
	fmt.Print("Enter the value")
	fmt.Scan(&n)
	fmt.Print("Enter the string")
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		str = append(str, s)
	}
	fmt.Println("The string is", str)

	result := validateSequence((str))

	for i := range result {
		fmt.Println(result[i])
	}

	fmt.Println("Hello, 世界")
}
