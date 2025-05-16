package linkedlist

import "fmt"

type DoublyLinkedList[T comparable] struct {
	Head *Node[T]
	Size uint
}

// Display traverses and prints the nodes of the list
func (dll *DoublyLinkedList[T]) Display() {
	if dll.Head == nil {
		println("List is empty")
		return
	}
	
	fmt.Print("Start -> ")
	for curr := dll.Head; curr != nil; curr = curr.Next {
		fmt.Printf("%v -> ", curr)
	}
	fmt.Println("")
}

// Prepend adds a new node to the Head of the list
func (dll *DoublyLinkedList[T]) Prepend(value T) {
	node := &Node[T]{Data: value}
	if dll.Head == nil {
		dll.Head = node
	} else {
		node.Next = dll.Head
		dll.Head.Prev = node
		dll.Head = node
	}
	
	dll.Size++
}

// Append adds a new node to the list of the list
func (dll *DoublyLinkedList[T]) Append(value T) {
	node := &Node[T]{Data: value}
	if dll.Head == nil {
		dll.Head = node
	} else {
		curr := dll.Head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = node
		node.Prev = curr
	}
	
	
	dll.Size++
}

// InsertAfter add a node with 'newValue' after the node whose 'Data' matches 'targetValue'
func (dll *DoublyLinkedList[T]) InsertAfter(targetValue T, newValue T) bool {
	if dll.Head == nil {
		return false
	}
	
	for curr := dll.Head; curr != nil; curr = curr.Next {
		if curr.Data == targetValue {
			node := &Node[T]{Data: newValue, Prev: curr, Next: curr.Next}
			if curr.Next != nil {
				node.Next.Prev = node
			}
			curr.Next = node
			dll.Size++
			return true
		}
	}
	return false
}

// InsertBefore add a node with 'newValue' before the node whose 'Data' matches 'targetValue'
func (dll *DoublyLinkedList[T]) InsertBefore(targetValue T, newValue T) bool {
	if dll.Head == nil {
		return false
	} else if dll.Head.Data == targetValue {
		node := &Node[T]{Data: newValue, Next: dll.Head}
		dll.Head.Prev = node
		dll.Head = node
		dll.Size++
		return true
	} else {
		for prev, curr := dll.Head, dll.Head.Next; curr != nil; prev, curr = curr, curr.Next {
			if curr.Data == targetValue {
				node := &Node[T]{Data: newValue, Next: curr, Prev: prev}
				curr.Prev = node
				prev.Next = node
				dll.Size++
				return true
			}
		}
	}
	
	return false
}
