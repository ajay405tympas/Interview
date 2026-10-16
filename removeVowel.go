// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func removeVowel(str string) {
	tempstr := ""
	for _, val := range str {
		switch val {
		case 'a', 'e', 'i', 'o', 'u':
			{
				continue
			}
		}
		tempstr = tempstr + string(val)

	}
	fmt.Println(tempstr)
}

func main() {

	str := "HelloWorld"
	removeVowel(str)
	fmt.Println("Hello, 世界")
}
