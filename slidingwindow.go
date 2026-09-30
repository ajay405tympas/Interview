// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func slidingWindow(arr []int, val int) []int {

	var resultArr []int
	fmt.Println("result intial", resultArr)
	for i := 0; i < len(arr)-val+1; i++ {
		s := arr[i : val+i]
		max := s[0]
		for j := 0; j < val; j++ {
			if s[j] > max {
				max = s[j]
				fmt.Println("max is \n", max)
			}
		}
		resultArr = append(resultArr, max)
	}
	return resultArr
}

func main() {

	arr := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 4

	result := slidingWindow(arr, k)
	fmt.Println("Result", result)

	fmt.Println("Hello, 世界")
}
