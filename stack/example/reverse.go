package stack

type node[T comparable] struct {
	elem T
	next *node[T]
}

type LinkedStack[T comparable] struct {
	top *node[T]
}

func NewLinkedStack[T comparable](capacity int) *LinkedStack[T] {
	return &LinkedStack[T]{top: nil}
}

func (stack *LinkedStack[T]) Push(value T) {
	stack.top = &node[T]{next: stack.top, elem: value}
}

func (stack *LinkedStack[T]) Pop() *node[T] {
	if stack.top == nil {
		return nil
	}
	
	value := &node[T]{elem: stack.top.elem}
	stack.top = stack.top.next
	return value
}

func (stack *LinkedStack[T]) IsEmpty() bool {
	return stack.top == nil
}

func ReverseArray[T comparable](array []T) []T {
	if array == nil || len(array) == 0 {
		return []T{}
	}
	
	stack := NewLinkedStack[T](len(array))
	for _, v := range array {
		stack.Push(v)
	}
		
	reversed := make([]T, len(array))
	index := 0
	for !stack.IsEmpty() {
		if v := stack.Pop(); v != nil {
			reversed[index] = v.elem
		}
		index++
	}
	
	return reversed
}

