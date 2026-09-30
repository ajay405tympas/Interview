// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type node struct {
	data int
	next *node
}

type Linkedlist struct {
	head *node
}

func (l *Linkedlist) insertEnd(val int) {
	newNode := &node{data: val}

	if l.head == nil {
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

func (l *Linkedlist) printList() {
	temp := l.head
	for temp != nil {
		fmt.Println("List val:", temp.data)
		temp = temp.next
	}
}

func (l *Linkedlist) deleteList(pos int) {

	prev := l.head

	if pos == 1 {
		prev = prev.next
		l.head = prev
		return
	}

	for i := 1; i < pos-1; i++ {
		fmt.Println("data\n", prev.data)
		prev = prev.next
	}
	fmt.Println("data--\n", prev.data)
	curr := prev.next
	prev.next = curr.next
	curr.next = nil
}

func main() {

	list := Linkedlist{}

	list.insertEnd(10)
	list.insertEnd(20)
	list.insertEnd(30)
	list.insertEnd(40)

	list.printList()

	list.deleteList(1)

	list.printList()

	fmt.Println("Hello, 世界")
}
