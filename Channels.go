// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan string)

	go func() {
		fmt.Println("In Routine")
		channel <- "ajay"
		fmt.Println("just now")
	}()

	time.Sleep(100)
	fmt.Println("Receive outside")
	value := <-channel
	fmt.Println(value)
	fmt.Println("Hello, 世界")
}
