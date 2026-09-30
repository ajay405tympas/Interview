// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var cond = sync.NewCond(&mutex)
var ready = false

func consumer(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("In Consumer")
	mutex.Lock()
	for !ready {
		fmt.Println("Consumer is waiting")
		cond.Wait()
	}
	fmt.Println("Consumer recieved signal")
	mutex.Unlock()

}
func producer(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("In Producer")
	mutex.Lock()
	ready = true
	fmt.Println("Producer signalling")
	cond.Signal()
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go consumer(&wg)
	go producer(&wg)

	wg.Wait()

	fmt.Println("Hello, 世界")
}
