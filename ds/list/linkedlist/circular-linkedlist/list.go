package linkedlist

import "fmt"

// Node represent each node within a linked list
type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

// CircularLinkedList is type of linked list, where the last node points to the first
// node forming a circular or looped linkage.
type CircularLinkedList[T comparable] struct {
	Head *Node[T]
	Size uint
}

// Display iterates and prints the nodes of the linked list
func (cll *CircularLinkedList[T]) Display() {
	if cll.Head == nil {
		fmt.Println("List is empty")
	}
	
	fmt.Println("Start -> ")
	
	curr := cll.Head
	for {
		fmt.Printf("%v -> ", curr)
		curr = curr.Next
		if curr == cll.Head {
			break
		}
	}
}
