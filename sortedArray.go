// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func sortedArray(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{3, 5, 6, 2, 4, 8}
	result := sortedArray(arr)
	fmt.Println("Sorted Array:", result)
	fmt.Println("Hello, 世界")
}
