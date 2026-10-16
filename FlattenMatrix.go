// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	list := [][]int{{1, 2}, {3, 4}}

	fmt.Println(list)
	newlist := []int{}

	for _, val := range list {
		for _, v := range val {
			newlist = append(newlist, v)
		}
	}
	fmt.Println(newlist)

	fmt.Println("Hello, 世界")
}
