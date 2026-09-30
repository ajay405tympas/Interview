// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func workerNode1(done chan bool) {
	fmt.Println("InWorkerNode1")
	done <- true

}

func workerNode2(done chan bool) {
	fmt.Println("InWorkerNode2")
	fmt.Println("testing", <-done)
}

func main() {
	x := make(chan bool)
	go workerNode1(x)

	//fmt.Println("InWorkerNode2")
	//fmt.Println("testing", <-x)

	workerNode2(x)
	fmt.Println("Hello, 世界")
}
