package stack

type node[T any] struct {
	data T
	next *node[T]
}

type StackLinkedList[T any] struct {
	top *node[T]
	size uint
}

// NewStackLinkedList initializes and returns a linkedlist-based
// Stack ADT
func NewLinkedList[T any]() *StackLinkedList[T] {
	return &StackLinkedList[T]{
		top: nil,
		size: 0,
	}
}

// Push adds a new node to the top of the stack
func (stack *StackLinkedList[T]) Push(value T) {
	stack.top = &node[T]{next: stack.top, data: value}
	stack.size++
}

// Pop removes and returns the top node of the stack.
func (stack *StackLinkedList[T]) Pop() *node[T] { 
	if stack.top == nil {
		return nil
	}
	
	node := stack.top
	stack.top = stack.top.next
	node.next = nil
	stack.size--
	return node
}

func (stack *StackLinkedList[T]) Peek() *node[T] { 
	if stack.top == nil {
		return nil
	}
	
	nCopy := new(node[T])
	nCopy.data = stack.top.data
	nCopy.next = stack.top.next
	
	return nCopy
}