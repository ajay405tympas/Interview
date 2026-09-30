// You can edit this code!
// Click here and start typing.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func workliveTask(ctx context.Context) {
	select {
	case <-Done():
		log.Println("Task cancelled:", ctx.Err())
		return
	default:
		
	}
}

func exyfunc(w http.ResponseWriter, r *http.request) {
	fmt.Fprintf(w, "Hello World this is exyfunc")
	
	ctx, cancel = context.withCancel(r.context())
	go workliveTask(ctx)
	
	time.Sleep(5 * time.Second)

	fmt.Fprintln(w, "Request completed")

}

func main() {

	var wg sync.WaitGroup

	http.handleFunc("/exy", exyfunc)

	fmt.Println("Server is running")
	http.ListenAndServe(":8080", nil)
	fmt.Println("Hello, 世界")
}
