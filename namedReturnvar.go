package main

import "fmt"

func tricky() (result int) {
	defer func() {
		result++
	}()
	return 5
}

func tricky2() (result int) {
	result = 5
	defer func() { result++ }()
	return // implicit return of named variable
}

func tricky3() int {
	defer func() { fmt.Println("defer runs") }()
	return 5 // no named return variable, so defer can't modify the return value
}

func main() {
	fmt.Println(tricky())
	fmt.Println(tricky2())
	fmt.Println(tricky3())
}
