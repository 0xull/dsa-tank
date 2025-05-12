package linkedlist

import "fmt"

type Node[T any] struct {
	Data T
	Next *Node[T]
}

type SinglyLinkedList[T any] struct {
	Head *Node[T]

	count uint
}

// Display iterates and prints the Data content of each node.
func (sll *SinglyLinkedList[T]) Display() {
	curr := sll.Head
	if curr == nil {
		fmt.Println("List is empty")
		return
	}
	
	fmt.Print("Start -> ")
	for curr != nil {
		fmt.Printf("%+v -> ", curr.Data)
		curr = curr.Next
	}
	fmt.Println("nil")
}

// 
