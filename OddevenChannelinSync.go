// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func printOdd(arr []int, och chan int, wg *sync.WaitGroup) {

	for i := range arr {
		if i%2 != 0 {
			och <- i
		}
	}
	close(och)
}

func printEven(arr []int, ech chan int, wg *sync.WaitGroup) {

	for i := range arr {
		if i%2 == 0 {
			ech <- i
		}
	}
	close(ech)
}

func main() {
	var wg sync.WaitGroup
	oddch := make(chan int)
	evench := make(chan int)

	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	wg.Add(1)
	go printOdd(arr, oddch, &wg)

	wg.Add(1)
	go printEven(arr, evench, &wg)

	for i := range arr {
		if i%2 == 0 {
			fmt.Println("even:", <-evench)
		} else {
			fmt.Println("even:", <-oddch)
		}
	}

	wg.Wait()

	fmt.Println("Hello, 世界")
}
