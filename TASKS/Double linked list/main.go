package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoublyLinkedList) Append(data int) {
	newNode := &Node{data: data}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	}
}

func (dll *DoublyLinkedList) Prepend(data int) {
	newNode := &Node{data: data}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	}
}

func (dll *DoublyLinkedList) Deappend() {
	if dll.tail == nil {
		return
	}

	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
		return
	}

	dll.tail = dll.tail.prev
	dll.tail.next = nil
}

func (dll *DoublyLinkedList) Defront() {
	if dll.head == nil {
		return
	}

	if dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
		return
	}

	dll.head = dll.head.next
	dll.head.prev = nil
}

func (dll *DoublyLinkedList) Display() {
	current := dll.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}

func (dll *DoublyLinkedList) Displayback() {
	current := dll.tail
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.prev
	}
	fmt.Println()
}

func main() {
	dll := DoublyLinkedList{}
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	fmt.Println("Doubly Linked List after Appending:")
	dll.Displayback()

	dll.Prepend(0)
	fmt.Println("Doubly Linked List after Prepending:")
	dll.Display()

	dll.Defront()
	fmt.Println("Doubly Linked List after Deleting Front:")
	dll.Display()

	dll.Deappend()
	fmt.Println("Doubly Linked List after Deappending:")
	dll.Display() 
}
