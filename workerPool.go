// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func worker(id int, result chan int, jobch chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	for i := range jobch {
		fmt.Printf("Job i-%d started by %d worker\n", i, id)
		result <- i * 2
		fmt.Printf("Job i-%d completed by %d worker\n", i, id)
	}

}

func main() {

	jobs := 10
	workers := 3

	var wg sync.WaitGroup

	jobsch := make(chan int, jobs)
	result := make(chan int, jobs)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(i, result, jobsch, &wg)
	}

	for j := 0; j < jobs; j++ {
		jobsch <- j
	}

	close(jobsch)
	wg.Wait() // wait for workers to finish
	close(result)

	for r := range result {
		fmt.Println("results:", r)
	}

	fmt.Println("Hello, 世界")
}
