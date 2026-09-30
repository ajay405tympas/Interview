// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"slices"
)

func main() {

	var count int
	nums := []int{100, 4, 200, 1, 3, 2}
	slices.Sort(nums)
	fmt.Println("sorted Array", nums)

	temp := nums[0]
	count = 1
	for i := 1; i < len(nums); i++ {
		temp = temp + 1
		if nums[i] == temp {
			count = count + 1
		}
	}
	fmt.Println("Consecutive array len", count)
	fmt.Println("Hello, 世界")
}
