package main

import (
	"fmt"
	"net/http"
)

func exHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func main() {
	http.HandleFunc("/ex", exHandler)

	fmt.Println("Server is listening on 8080")

	http.ListenAndServe(":8080", nil)
}
