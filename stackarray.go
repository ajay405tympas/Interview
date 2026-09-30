// You can edit this code!
// Click here and start typing.
package main

import "fmt"

var stack []int

func push(val int) {
	stack = append(stack, val)
}

func pop() int {
	var pop int
	if len(stack) == 0 {
		fmt.Println("Stack empty")
	} else {
		pop = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return pop
}

func peek() int {
	return stack[len(stack)-1]
}

func main() {

	push(10)
	push(20)

	push(30)

	fmt.Println("Pop Element", pop())

	fmt.Println("Peek Element", peek())

	fmt.Println("Hello, 世界")
}
