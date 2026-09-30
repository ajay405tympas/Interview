// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func getMinSubArray(b []int) int {
	min := b[0]
	for p := 1; p < len(b); p++ {
		if min > b[p] {
			min = b[p]
		}
	}
	//fmt.Println("min", min)
	return min
}

func getMaxSubArray(b []int) int {
	max := b[0]
	for p := 1; p < len(b); p++ {
		if max < b[p] {
			max = b[p]
		}
	}
	//fmt.Println("max", max)
	return max
}

func minMaxArray(a []int) []int {

	maxArr := make([]int, len(a))
	maxVal := a[0]

	var minVal = make([]int, len(a))

	for k := 1; k < len(a); k++ {
		if maxVal < a[k] {
			maxVal = a[k]
		}
	}
	maxArr[0] = maxVal

	for k := 2; k <= len(a); k++ {
		for i := 0; i < len(a)-k+1; i++ {
			var temp = make([]int, k)
			temp = append(a[i : k+i])
			minVal[i] = getMinSubArray(temp)
		}

		//fmt.Println("minvalArray", minVal)
		maxArr[k-1] = getMaxSubArray(minVal)
		fmt.Println("maxArr", maxArr)
	}
	return maxArr
}

func main() {
	arr := []int{6, 3, 5, 1, 12}

	fmt.Println("Final result", minMaxArray(arr))
	fmt.Println("Hello, 世界")
}
