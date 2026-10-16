// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	list := []string{"apple", "bananan", "Nework", "Raghuvendr"}
	longwordindex := 0
	count := 0
	for i, str := range list {
		l := len(str)
		if count < l {
			count = l
			longwordindex = i
		}
	}
	fmt.Println("largest word", list[longwordindex])
	fmt.Println("Hello, 世界")
}
