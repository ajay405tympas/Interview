// // You can edit this code!
// // Click here and start typing.
 package main

 import "fmt"

 func counter() func() int {
 	count := 0
 	return func() int {
 		count++
 		return count
 	}
 }

 func main() {
 	c1 := counter()

 	fmt.Println("c1-->", c1())
 	fmt.Println("c1-->", c1())
 	fmt.Println("c1-->", c1())

 	fmt.Println("Hello, 世界")
 }
