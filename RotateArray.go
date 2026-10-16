// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func rotateArr(a []int, k int) []int {
	n := len(a)
	k = k % n
	return append(a[n-k:], a[:n-k]...)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	k := 1
	fmt.Println("Before Array", arr)

	fmt.Println("After Array", rotateArr(arr, k))

	fmt.Println("Hello, 世界")
}
