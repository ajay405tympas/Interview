// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type stack struct {
	data int
	next *stack
}

type stackImpl struct {
	top *stack
}

func (s *stackImpl) push(val int) {
	newNode := &stack{data: val}
	if s.top == nil {
		s.top = newNode
		newNode.next = nil
		return
	}

	newNode.next = s.top
	s.top = newNode

}

func (s *stackImpl) pop() *stack {

	if s.top == nil {
		fmt.Println("stack is under flow")
		return nil
	}
	temp := s.top
	s.top = s.top.next

	return temp

}

func (s *stackImpl) printStackTop() {
	fmt.Println("Top Poiniting to:", s.top.data)
}

func main() {

	stack := stackImpl{}
	stack.push(10)
	stack.push(30)
	stack.push(40)
	stack.push(50)

	stack.printStackTop()

	stack.pop()

	stack.printStackTop()

	fmt.Println("Hello, 世界")
}
