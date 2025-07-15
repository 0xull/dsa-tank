package stack

type StackArray[T comparable] struct {
	items []T
}

func (stack *StackArray[T]) Push(value T) {
	stack.items = append(stack.items, value)
}

func (stack *StackArray[T]) Pop() T {
	top := len(stack.items) - 1
	value := stack.items[top]
	stack.items = stack.items[:top]
	return value
}

func (stack *StackArray[T]) Peek() T {
	return stack.items[len(stack.items)-1]
}

func (stack *StackArray[T]) IsEmpty() bool {
	return len(stack.items) == 0
} 

func (stack *StackArray[T]) Size() int {
	return len(stack.items)
}