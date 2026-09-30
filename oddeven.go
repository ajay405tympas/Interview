// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func printOdd(odd chan int, even chan int, max int, wg *sync.WaitGroup) {

	defer wg.Done()
	for val := range odd {
		if val > max {
			close(even)
			return
		}

		fmt.Println("Odd--", val)
		even <- val + 1

	}

}

func printEven(odd chan int, even chan int, max int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range even {
		if val > max {
			close(odd)
			return
		}

		fmt.Println("Even--", val)
		odd <- val + 1

	}

}

func main() {

	max := 100
	var wg sync.WaitGroup
	oddch := make(chan int)
	evench := make(chan int)

	wg.Add(2)
	go printOdd(oddch, evench, max, &wg)
	go printEven(oddch, evench, max, &wg)

	oddch <- 1

	wg.Wait()
	fmt.Println("Hello, 世界")
}
