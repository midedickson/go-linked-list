package main

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

func CreateNode(value interface{}) *Node {
	return &Node{Value: value}
}

// List respresents the implementation of a linked list
type List struct {
	Head *Node
}

func (l *List) init(value interface{}) {
	// initialize the linked list with the head node
	l.Head = CreateNode(value)
}

func (l *List) Push(value interface{}) {
	prevNode := l.Head
	currNode := l.Head.Next
	for {
		if currNode == nil {
			newNode := CreateNode(value)
			prevNode.Next = newNode
			break
		} else {
			prevNode = currNode
			currNode = prevNode.Next
		}
	}
}

func (l *List) Pop() interface{} {
	prevNode := l.Head
	currNode := l.Head.Next
	for {
		if currNode == nil {
			prevNode.Next = currNode.Next
			return currNode.Value
		} else {
			prevNode = currNode
			currNode = prevNode.Next
		}
	}
}

func (l *List) Reverse() {

	var prevNode *Node
	currNode := l.Head
	for {
		if currNode == nil {
			break
		} else {
			dummyNode := currNode.Next
			currNode.Next = prevNode
			prevNode = currNode
			currNode = dummyNode
		}
	}

	l.Head = prevNode

}

func (l *List) Print() {
	var prevNode *Node
	currNode := l.Head
	for {
		if currNode == nil {
			fmt.Println("End of List")
			break
		} else {
			fmt.Printf("Value: %v\n", currNode.Value)
			prevNode = currNode
			currNode = prevNode.Next
		}
	}
}

func CreateList(value interface{}) *List {
	newList := &List{}
	newList.init(value)
	return newList
}

func main() {
	l := CreateList(3)
	l.Print()

	l.Push(5)
	l.Push(6)
	l.Push(7)
	l.Push(8)
	l.Push(9)

	l.Print()

	l.Reverse()
	l.Print()

}
