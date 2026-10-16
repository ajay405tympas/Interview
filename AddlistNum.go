// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	list := []int{1, 2, 8, 15, 5, 7}
	//fmt.Scanf(&target)
	target := 9
	seen := make(map[int]int)

	for i, val := range list {
		stillneed := target - val

		id, found := seen[stillneed]
		if found {
			fmt.Println(id)
			fmt.Println(" ", i)
			continue
		} else {
			seen[val] = i
		}

	}

	fmt.Println("Hello, 世界")
}
