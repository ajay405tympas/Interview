// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func maxProfitStock(arr []int) int {
	minPrice := arr[0]
	maxProfit := 0
	for i := 1; i < len(arr); i++ {
		if minPrice > arr[i] {
			minPrice = arr[i]
		} else {
			profit := arr[i] - minPrice
			if profit > maxProfit {
				maxProfit = profit
			}

		}
	}
	return maxProfit
}

func main() {

	arr := []int{9, 2, 5, 3, 6, 8, 9}
	fmt.Println("Max profit after:", maxProfitStock(arr))
	fmt.Println("Hello, 世界")
}
