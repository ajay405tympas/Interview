package main

import "fmt"

func moveZeros(arr []int) []int {
	j := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
	return arr

}

func main() {

	arr := []int{2, 1, 3, 0, 12, 4, 0}

	result := moveZeros(arr)

	fmt.Println("Modified Array", result)

}
