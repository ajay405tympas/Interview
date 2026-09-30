// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

func highestFrquency(arr []int) (string, int) {
	highFreqmap := make(map[string]int)

	for i := 0; i < len(arr); i++ {
		keyValue := strconv.Itoa(arr[i])
		_, exists := highFreqmap[keyValue]
		if !exists {
			highFreqmap[keyValue] = 1
		} else {
			val, _ := highFreqmap[keyValue]
			highFreqmap[keyValue] = val + 1
		}
	}
	maxValue := 0
	maxKey := ""
	fmt.Println("TotalMap", highFreqmap)
	for key, value := range highFreqmap {
		if value > maxValue {
			maxValue = value
			maxKey = key
		}
	}
	return maxKey, maxValue
}

func main() {
	arr := []int{1, 3, 2, 3, 4, 3, 2, 2, 4, 1, 2}

	key, value := highestFrquency(arr)

	fmt.Println("Highest occurence:", key, value)

	fmt.Println("Hello, 世界")
}
