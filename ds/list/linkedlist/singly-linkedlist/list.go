package linkedlist

import "fmt"

// Node represent each node in a linked list.
type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

// SinglyLinkedList stored data in linear unidirectional 
// chains.
type SinglyLinkedList[T comparable] struct {
	Head *Node[T]

	count uint
}

// Display iterates and prints the Data content of each node.
func (sll *SinglyLinkedList[T]) Display() {
	if sll.Head == nil {
		fmt.Println("List is empty")
		return
	}
	
	curr := sll.Head
	fmt.Print("Start -> ")
	for curr != nil {
		fmt.Printf("%+v -> ", curr.Data)
		curr = curr.Next
	}
	fmt.Println("nil")
}

// Search iterates and compares the data at each node for a match
// against 'value', returning the address.
func (sll *SinglyLinkedList[T]) Search(value T) *Node[T] {
	curr := sll.Head
	for curr != nil {
		if curr.Data == value {
			return curr
		}
		curr = curr.Next
	}
	
	return nil
}

// Prepend attaches a new node with 'value' to the beginning
// of the list.
func (sll *SinglyLinkedList[T]) Prepend(value T) {
	sll.Head = &Node[T]{Data: value, Next: sll.Head}
	sll.count++
}


