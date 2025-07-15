package linkedlist

import "fmt"

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
		return
	}
	
	fmt.Print("Start -> ")
	
	curr := cll.Head
	for {
		fmt.Printf("%v -> ", curr.Data)
		curr = curr.Next
		if curr == cll.Head {
			fmt.Print("(Head)\n")
			break
		}
	}
}

// Prepend attaches a node with 'value' as the new 'Head' of
// the list.
func (cll *CircularLinkedList[T]) Prepend(value T) {
	node :=  &Node[T]{Data: value}
	if cll.Head == nil {
		cll.Head = node
		cll.Head.Next = cll.Head
	} else {
		for curr := cll.Head;; curr = curr.Next {
			if curr.Next == cll.Head {
				node.Next = cll.Head
				cll.Head = node
				curr.Next = cll.Head
				break
			}
		}
	}
	cll.Size++
}

// Append adds a node with 'value' data to the end of the list.
func (cll *CircularLinkedList[T]) Append(value T) {
	node := &Node[T]{Data: value}
	if cll.Head == nil {
		cll.Head = node
		node.Next = cll.Head
	} else {
		for curr := cll.Head;; curr = curr.Next {
			if curr.Next == cll.Head {
				node.Next = cll.Head
				curr.Next = node
				break
			}
		}
	}
	cll.Size++
}

// DeleteHead removes and return the current 'Head' node from
// the list.
func (cll *CircularLinkedList[T]) DeleteHead() *Node[T] {
	if cll.Head == nil {
		return nil
	}
	
	node := cll.Head
	if cll.Size == 1 {
		cll.Head = nil
		cll.Size--
		return node
	}
	
	for curr := cll.Head.Next;; curr = curr.Next {
		if curr.Next == cll.Head {
			curr.Next = cll.Head.Next
			cll.Head = cll.Head.Next
			cll.Size--
			return node
		}
	}
}

// DeleteLast removes and returns the last node in list
func (cll *CircularLinkedList[T]) DeleteLast() *Node[T] {
	if cll.Head == nil {
		return nil
	}
	
	var node *Node[T]
	if cll.Head == cll.Head.Next {
		node = cll.Head
		cll.Head = nil
	} else {
		for prev, curr := cll.Head, cll.Head.Next;; prev, curr = curr, curr.Next {
			if curr.Next == cll.Head {
				node = curr
				prev.Next = cll.Head
				break
			}
		}
	}
	cll.Size--
	return node
}

// Delete removes and returns a node that its data matches the 'targetValue'
func (cll *CircularLinkedList[T]) Delete(targetValue T) *Node[T] {
	if cll.Head != nil {
		var node *Node[T]
		if cll.Head.Data == targetValue {
			node = cll.Head
			if cll.Head == cll.Head.Next {
				cll.Head = nil
				cll.Size--
				return node
			}
			
			for curr := cll.Head.Next;; curr = curr.Next {
				if curr.Next == cll.Head {
					curr.Next = cll.Head.Next
					cll.Head = cll.Head.Next
					cll.Size--
					return node
				}
			}
		}
		
		for prev, curr := cll.Head, cll.Head.Next; curr != cll.Head; prev, curr = curr, curr.Next {
			if curr.Data == targetValue {
				node = curr
				prev.Next = curr.Next
				cll.Size--
				return node
			}
		}
	}
	return nil
}