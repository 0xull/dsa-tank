package stack

type StackArray[T comparable] struct {
	items []T
	top int
	capacity int
}

func NewStackArray[T comparable](capacity int) *StackArray[T] {
	if capacity <= 0 {
		capacity = 0
	}
	
	return &StackArray[T]{
		items: make([]T, capacity),
		top: -1,
		capacity: capacity,
	}
}

func (stack *StackArray[T]) Push(value T) {
	if stack.top == (stack.capacity -1) {
		return
	}
	
	stack.top++
	stack.items[stack.top] = value
}

func (stack *StackArray[T]) Pop() T {
	var value T
	if stack.top == -1 {
		return value
	}
	
	value = stack.items[stack.top]
	stack.items[stack.top] = *new(T)
	stack.top--
	return value
}

func (stack *StackArray[T]) Peek() T {
	if stack.top == -1 {
		return *new(T)
	}
	
	return stack.items[stack.top]
}
