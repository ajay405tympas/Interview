// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func maxresultAndSubArray(arr []int) ([]int, int) {
	maxSum := arr[0]
	currentSum := arr[0]
	start, end := 0, 0
	tempStart := 0

	for i := 1; i < len(arr); i++ {
		if currentSum+arr[i] > arr[i] {
			currentSum = currentSum + arr[i]
		} else {
			currentSum = arr[i]
			tempStart = i
		}

		if currentSum > maxSum {
			maxSum = currentSum
			start = tempStart
			end = i
		}
	}
	fmt.Println("StartArray: endArray", start, tempStart, end)
	return arr[start : end+1], maxSum
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 6, 2, 3, 2}

	subarr, result := maxresultAndSubArray(arr)

	fmt.Println("SubArray\n", subarr)
	fmt.Println("Result\n", result)

	fmt.Println("Hello, 世界")
}
