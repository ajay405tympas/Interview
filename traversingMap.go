// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	m := map[string][]int{
		"A": []int{8, 2, 3},
		"B": []int{2},
		"C": []int{3},
	}

	for i, k := range m {
		fmt.Println("i-k", i, k)
		for _, l := range k {
			fmt.Println("l", l)
		}
	}
	fmt.Println("Hello, 世界")
}


// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type card struct {
	name   string
	cardno int
}

func main() {

	/*graph := map[string][]int[]int  {
		//"A": {
		//	{ name = "Ax",cardno: 123},
		//	{ name = "Ay",cardno: 324},
		//     },
		//"B": {
	//		{name = "Bx",cardno: 456},
	//		{name = "By",cardno: 457},
	//	     },
	//	"C": {
	//		{name = "Cx",cardno: 446},
	//		{name = "Cy",cardno: 17},
	//	     },
	//	}*/

	//graph := map[string][][]int
	
	A:{
	    [1 2 3] [3 4 5] 
	},
	B: 
	
	
	fmt.Println("Hello, 世界")
}
