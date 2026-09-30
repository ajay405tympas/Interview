// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedlist struct {
	head *node
}

func (l *linkedlist) swapNodes(PosA int, PosB int) {
	temp := l.head

	A := &node{}
	B := &node{}

	if PosA == 0 {
		A = temp
	} else {
		for i := 0; i < PosA; i++ {
			temp = temp.next
		}
		A = temp
	}

	temp = l.head
	if PosB == 0 {
		B = temp
	} else {
		for i := 0; i < PosB; i++ {
			temp = temp.next
		}
		B = temp
	}

	if PosA == 0 || PosB == 0 {
		A.next = B.next
		B.next = A
		l.head = B
	} else {
		prev := l.head
		for j := 0; j < PosA-1; j++ {
			prev = prev.next
		}
		A.next = B.next
		B.next = A
		prev.next = B
	}
}

func (l *linkedlist) insertNode(val int) {
	newNode := &node{data: val}

	if l.head == nil {
		newNode.next = nil
		l.head = newNode
		return
	}
	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}

	temp.next = newNode
	newNode.next = nil
}

func (l *linkedlist) printNode() {
	temp := l.head
	for temp != nil {
		fmt.Println(temp.data)
		temp = temp.next
	}
}

func (l *linkedlist) deleteNode(pos int) {
	temp := l.head

	if pos == 1 {
		//deleting head
		temp = temp.next
		l.head = temp
		return
	}

	for i := 0; i < pos; i++ {
		temp = temp.next
	}
	prev := temp
	curr := temp.next
	prev.next = curr.next
	curr.next = nil
}

func (l *linkedlist) swapNodeMethod() {
	temp := l.head
	count := 0
	for temp != nil {
		count++
		temp = temp.next
	}

	for i := 0; i < count/2; i++ {
		PosA := 2 * i
		PosB := 2*i + 1
		fmt.Println(" swapping pairs :", PosA, PosB)
		l.swapNodes(PosA, PosB)
	}
}

func main() {

	list := linkedlist{}

	list.insertNode(10)
	list.insertNode(20)
	list.insertNode(30)
	list.insertNode(40)
	list.insertNode(50)
	list.insertNode(60)
	list.insertNode(70)

	list.printNode()

	list.swapNodeMethod()

	list.printNode()

	fmt.Println("Hello, 世界")
}
