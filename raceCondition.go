// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

var count int

func countermsg(wg *sync.WaitGroup) {
	defer wg.Done()
	count++
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go countermsg(&wg)
	}

	wg.Wait()
	fmt.Println("Final val:", count)
	fmt.Println("Hello, 世界")
}
