// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{1, 1, 1, 1, 4, 4, 5, 5, 5, 2, 3, 0, 2, 3}
	freqmap := make(map[int]int)

	for _, val := range nums {
		_, found := freqmap[val]
		if found {
			freqmap[val] = freqmap[val] + 1
		} else {
			freqmap[val] = 1
		}
	}

	fmt.Println("UnsortedMap:", freqmap)

	var keys []int

	for k := range freqmap {
		keys = append(keys, k)
	}

	slices.SortFunc(keys, func(a, b int) int {
		if freqmap[a] == freqmap[b] {
			return b - a
		}
		return freqmap[a] - freqmap[b]

	})

	fmt.Println(keys)

	fmt.Println("Hello, 世界")
}
