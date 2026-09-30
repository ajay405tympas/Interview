// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type node struct{
	data int
	next *node
}
type linkedList struct{
	head *node
}

func (l *linkedList)insertNode(data int){
	newNode = &node{data:data}
	if head == nil{
		l.head = newNode
		newNode.next = nil
		return
	}
	temp :=l.head
	for temp.next != nil {
	
		temp = temp.next
	}
	temp = newNode
	newNode.next = nil
}

func (l *linkedList)replaceNode(val int){
	temp := head
	counter := 0
	for temp.next != nil{
		if val >= temp.data{
			counter++
		}
	}
}

func (l *linkedList)printNode(){
}

func main() {
	arr := []int{1,1,1,2,2,3,6,4,3,2,5}
	k:=3
	list = linkedList{}
	//setting up initial k nodes
	for i:=0;i<k;i++ {
	   list.insertNode(arr[i])
	}
	
	for(i:=k;i<len(arr);i++{
		replaceNode(arr[k])
	}
	fmt.Println("Hello, 世界")
}