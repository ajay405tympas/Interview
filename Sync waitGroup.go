// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func goWorkerNode(id int, wG *sync.WaitGroup) {
	defer wG.Done()
	fmt.Println("in WorkerNode", id)
}

func main() {
	var Wg sync.WaitGroup
	for i := 0; i <= 3; i++ {
		Wg.Add(1)
		go goWorkerNode(i, &Wg)
	}

	Wg.Wait()
	fmt.Println("all goworkenodes")
	fmt.Println("Hello, 世界")
}
