// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func removeElementFromArr(arr []int, pos int) []int {

	fmt.Println("Before removing", arr, pos)
	for i, _ := range arr {
		if i == pos {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	fmt.Println("after removing", arr)
	return arr
}

func getLargestElement(arr []int, val int) []int {

	resultArr := make([]int, val)
	for i := 0; i < val; i++ {
		max := arr[0]
		position := 0
		fmt.Println("initialMax - length ", max, len(arr))
		for j := 1; j < len(arr); j++ {
			if max < arr[j] {
				max = arr[j]
				position = j
				fmt.Println("max", max, position)
			}
		}
		arr = removeElementFromArr(arr, position)
		resultArr[i] = max
		fmt.Println("resultArr", resultArr[i])
	}
	return resultArr
}

func main() {
	arr := []int{1, 3, 2, 3, 4, 3, 5, 7, 8, 2}
	k := 4
	fmt.Println("Highest numbers in Array", getLargestElement(arr, k))

	fmt.Println("Hello, 世界")
}
