package linkedlist

import "fmt"

// SinglyLinkedList stored data in linear unidirectional
// chains.
type SinglyLinkedList[T comparable] struct {
	Head *Node[T]

	Count uint
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
	sll.Count++
}

// Append attaches a new node with 'value' to the end of the list.
// If frequent appends are major use case, you can maintain an additional
// node, say Tail, pointing to the last node in the list.
func (sll *SinglyLinkedList[T]) Append(value T) {
	node := &Node[T]{Data: value}
	if sll.Head == nil {
		sll.Head = node
	} else {
		for curr := sll.Head;; curr = curr.Next {
			if curr.Next == nil {
				curr.Next = node
				break
			}
		}
	}
	sll.Count++
}

// InsertAfter adds a new node with 'newNodeData' after the node with specified
// 'targetValue'.
func (sll *SinglyLinkedList[T]) InsertAfter(targetValue T, newNodeData T) bool {
	for curr := sll.Head; curr != nil; curr = curr.Next {
		if curr.Data == targetValue {
			node := &Node[T]{Data: newNodeData, Next: curr.Next}
			curr.Next = node
			sll.Count++
			return true
		}
	}

	return false
}

// InsertBefore adds a node with 'newNodeData' before the node with 'targetValue'
func (sll *SinglyLinkedList[T]) InsertBefore(targetValue T, newNodeData T) bool {
	if sll.Head == nil {
		return false
	}

	node := &Node[T]{Data: newNodeData, Next: sll.Head}
	if sll.Head.Data == targetValue {
		sll.Head = node
		sll.Count++
		return true
	}

	for prev, curr := sll.Head, sll.Head.Next; curr != nil; prev, curr = curr, curr.Next {
		if curr.Data == targetValue {
			node.Next = prev.Next
			prev.Next = node
			sll.Count++
			return true
		}
	}

	return false
}

// DeleteHead removes and returns the current head's node
func (sll *SinglyLinkedList[T]) DeleteHead() *Node[T] {
	if sll.Head == nil {
		return nil
	}
	
	node := sll.Head
	sll.Head = sll.Head.Next
	sll.Count--
	return node
}

// DeleteTail removes and returns the current tail's node
func (sll *SinglyLinkedList[T]) DeleteTail() *Node[T] {
	if sll.Head == nil {
		return nil
	}
	
	if sll.Head.Next == nil {
		node := sll.Head
		sll.Head = nil
		sll.Count--
		return node
	}
	
	for prev, curr := sll.Head, sll.Head.Next;; prev, curr = curr, curr.Next {
		if curr.Next == nil {
			prev.Next = nil
			sll.Count--
			return curr
		}
	}
}

// Delete removes and returns a node with 'targetValue'
func (sll *SinglyLinkedList[T]) Delete(targetValue T) *Node[T] {
	if sll.Head == nil {
		return nil
	}
	
	if sll.Head.Data == targetValue {
		node := sll.Head
		sll.Head = sll.Head.Next
		sll.Count--
		return node
	}
	
	for prev, curr := sll.Head, sll.Head.Next; curr != nil; prev, curr = curr, curr.Next {
		if curr.Data == targetValue {
			prev.Next = curr.Next
			sll.Count--
			return curr
		}
	}
	
	return nil
}
