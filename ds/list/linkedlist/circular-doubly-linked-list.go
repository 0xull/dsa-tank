package linkedlist

import "fmt"

type CircularDoublyLinkedList[T comparable] struct {
	Head *Node[T]
	Size uint
}

// Display prints all nodes in the list
func (cdll *CircularDoublyLinkedList[T]) Display() {
	if cdll.Head == nil {
		fmt.Println("List is empty")
		return
	}
	
	curr := cdll.Head
	for curr.Next != cdll.Head {
		fmt.Printf("%v <-> ", curr)
		curr = curr.Next
	}
	fmt.Printf("%v <-> (Head)\n", curr)
}

// Prepend adds node to the beginning of the list
func (cdll *CircularDoublyLinkedList[T]) Prepend(value T) {
	node := &Node[T]{Data: value}
	if cdll.Head != nil {
		oldHead, tail := cdll.Head, cdll.Head.Prev 
		oldHead.Prev, tail.Next = node, node
		
		cdll.Head = node
		cdll.Head.Next = oldHead
		cdll.Head.Prev = tail
	} else {
		cdll.Head = node
		cdll.Head.Next = cdll.Head
		cdll.Head.Prev = cdll.Head
	}
	
	cdll.Size++
}

// Append adds a node to the ending of the list
func (cdll *CircularDoublyLinkedList[T]) Append(value T) {
	node := &Node[T]{Data: value}
	if cdll.Head == nil {
		cdll.Head = node
		cdll.Head.Next = cdll.Head
		cdll.Head.Prev = cdll.Head
	} else {
		node.Next = cdll.Head
		node.Prev = cdll.Head.Prev
		
		cdll.Head.Prev.Next = node
		cdll.Head.Prev = node
	}

	cdll.Size++
}
