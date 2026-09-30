// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
	"time"
)

func semPatternProgram(i int, limit chan string, wg *sync.WaitGroup) {

	limit <- ""

	defer func() { <-limit }()
	defer wg.Done()

	fmt.Printf("%d-started\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d-completd\n", i)
}

func main() {

	var wg sync.WaitGroup
	limiter := make(chan string, 4)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go semPatternProgram(i, limiter, &wg)
	}

	wg.Wait()
	close(limiter)
	fmt.Println("Hello, 世界")
}
