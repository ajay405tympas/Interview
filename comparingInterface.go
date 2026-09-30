// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	var a interface{} = []int{1, 2}
	var b interface{} = []int{1, 2}

	fmt.Println(a[0] == b[0])
}

/*package main

import "fmt"

func main() {
	var a interface{} = 10
	var b int = 10

	//i := a.(int)
	//j := b.(int)

	fmt.Println(a == b)
}*/
