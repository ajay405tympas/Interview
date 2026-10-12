 // You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func leakyworker(done chan int) {
	for {
		<-done
	}
}

func goTestRoutine(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Goroutines %d", runtime.NumGoroutine())
}

func main() {
	channel := make(chan int)

	for  {
		go leakyworker(channel)
	}
	http.HandleFunc("/goTestRoutine", goTestRoutine)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)

	//fmt.Println("Hello, 世界")
}
