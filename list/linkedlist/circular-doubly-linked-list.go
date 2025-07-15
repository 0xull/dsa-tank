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

// InsertBefore adds a node with 'value' before the node with 'targetValue'
func (cdll *CircularDoublyLinkedList[T]) InsertBefore(value T, targetValue T) bool {
	if cdll.Head == nil {
		return false
	}
	
	if cdll.Head.Data == targetValue {
		node := &Node[T]{Data: value, Next: cdll.Head, Prev: cdll.Head.Prev}
		cdll.Head.Prev.Next = node
		cdll.Head.Prev = node
		cdll.Head = node
		cdll.Size++
		return true
	} else {
		for curr, prev := cdll.Head.Next, cdll.Head; curr != cdll.Head; prev, curr = curr, curr.Next {
			if curr.Data == targetValue {
				node := &Node[T]{Data: value, Next: curr, Prev: prev}
				prev.Next = node
				curr.Prev = node
				cdll.Size++
				return true
			}
		}
	}
	
	return false
}

// InsertAfter adds a node with 'value' after a node with 'targetValue'.
func (cdll *CircularDoublyLinkedList[T]) InsertAfter(value T, targetValue T) bool {
	if cdll.Head == nil {
		return false
	}
	
	if cdll.Head.Data == targetValue {
		node := &Node[T]{Data: value, Next: cdll.Head.Next, Prev: cdll.Head}
		cdll.Head.Next = node
		cdll.Head.Next.Prev = node
		cdll.Size++
		return true
	} else {
		for curr := cdll.Head.Next; curr != cdll.Head; curr = curr.Next {
			if curr.Data == targetValue {
				node := &Node[T]{Data: value, Next: curr.Next, Prev: curr}
				curr.Next.Prev = node
				curr.Next = node
				cdll.Size++
				return true
			}
		}
	}
	return false
}

// DeleteHead removes and returns the current 'Head' node
func (cdll *CircularDoublyLinkedList[T]) DeleteHead() *Node[T] {
	if cdll.Head == nil {
		return nil
	}
	
	head := cdll.Head
	
	if head.Next == cdll.Head {
		cdll.Head = nil
	} else {
		cdll.Head.Next.Prev = cdll.Head.Prev
		cdll.Head.Prev.Next = cdll.Head.Next
		cdll.Head = cdll.Head.Next
	}
	
	cdll.Size--
	
	head.Prev = nil
	head.Next = nil
	return head
}

func (cdll *CircularDoublyLinkedList[T]) DeleteTail() *Node[T] {
	if cdll.Head == nil {
		return nil
	}
	
	tail := cdll.Head.Prev
	if tail == cdll.Head {
		cdll.Head = nil
	} else {
		tail.Prev.Next = cdll.Head
		cdll.Head.Prev = tail.Prev
	}
	
	cdll.Size--
	
	tail.Next = nil
	tail.Prev = nil
	return tail
}

func (cdll *CircularDoublyLinkedList[T]) Delete(targetValue T) *Node[T] {
	if cdll.Head == nil {
		return nil
	}
	
	if cdll.Head.Data == targetValue {
		node := cdll.Head
		if cdll.Head == cdll.Head.Next {
			cdll.Head = nil
			cdll.Size--
			return node
		} else {
			cdll.Head.Next.Prev = cdll.Head.Prev
			cdll.Head.Prev.Next = cdll.Head.Next
			cdll.Head = cdll.Head.Next
			
			cdll.Size--
			return node
		}
	}
	
	for prev, curr := cdll.Head, cdll.Head.Next; curr != cdll.Head; prev, curr = curr, curr.Next {
		if curr.Data == targetValue {
			node := curr
			prev.Next = curr.Next
			curr.Next.Prev = prev
			
			cdll.Size--
			node.Next = nil
			node.Prev = nil
			return node
		}
	}
	
	return nil
}
