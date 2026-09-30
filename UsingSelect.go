 // You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

func thread1(x chan string) {
	time.Sleep(1 * time.Second)
	x <- "Test1"

}

func thread2(x chan string) {
	time.Sleep(3 * time.Second)
	x <- "Test2"
}

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go thread1(channel1)
	go thread2(channel2)

	select {
	case msg1 := <-channel1:
		fmt.Println("channel 1", msg1)
		return
	case msg2 := <-channel2:
		fmt.Println("channel 2", msg2)
		return
		//default:
		//	fmt.Println("Astla Vistla Bye")

	}

	fmt.Println("Hello, 世界")
}
