// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	defer fmt.Println("done")
	defer fmt.Println("A")
	defer fmt.Println("B")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()

	panic("Mayday!!")
	fmt.Println("done")

}
