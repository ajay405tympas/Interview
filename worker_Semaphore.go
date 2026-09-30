// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, result chan int, jobch chan int, limit chan string, wg *sync.WaitGroup) {

	limit <- ""

	defer wg.Done()
	defer func() { <-limit }()

	for i := range jobch {
		fmt.Printf("Job i-%d started by %d worker\n", i, id)
		result <- i * 2
		time.Sleep(3 * time.Second)
		fmt.Printf("Job i-%d completed by %d worker\n", i, id)
	}

}

func main() {

	jobs := 10
	workers := 10

	var wg sync.WaitGroup

	limit := make(chan string, 2)
	jobsch := make(chan int, jobs)
	result := make(chan int, jobs)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(i, result, jobsch, limit, &wg)
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
