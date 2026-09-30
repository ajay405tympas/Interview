// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	runes := []rune("Ajay")
	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}

	fmt.Println(string(runes))
	fmt.Println("Hello, 世界")
}
