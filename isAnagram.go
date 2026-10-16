// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	strA := "silent"
	strB := "listen"

	strAmap := make(map[string]int)

	for _, ch := range strA {
		id, found := strAmap[string(ch)]
		if found {
			id = id + 1
			strAmap[string(ch)] = id
		} else {
			strAmap[string(ch)] = 0
		}
	}

	fmt.Println(strAmap)
	isAnagram := false
	for _, qch := range strB {
		_, found := strAmap[string(qch)]
		if found {
			isAnagram = true
			continue
		} else {
			isAnagram = false
		}
	}

	if isAnagram {
		fmt.Println("is Anagram")
	} else {
		fmt.Println("is not Anagram")
	}

	fmt.Println("Hello, 世界")
}
