package stack

import "fmt"

type MultiStack struct {
	items []any
	max int
	head int
	tail int
}

func NewMultiStack(capacity int) *MultiStack {
	if capacity <= 0 {
		capacity = 10
	}
	
	return &MultiStack{
		items: make([]any, capacity),
		max: capacity,
		head: -1,
		tail: capacity,
	}
}

func (stack *MultiStack) PushHead(value any) error {
	if stack.head + 1 == stack.tail {
		return fmt.Errorf("stack head overflow")
	}
	
	stack.head++
	stack.items[stack.head] = value
	return nil
}

func (stack *MultiStack) PushTail(value any) error {
	if stack.tail - 1 == stack.head {
		return fmt.Errorf("stack tail overflow")
	}
	
	stack.tail--
	stack.items[stack.tail] = value
	return nil
}

func (stack *MultiStack) PopHead() any {
	if stack.head <= -1 {
		return nil
	}
	
	value := stack.items[stack.head]
	stack.items[stack.head] = nil
	stack.head--
	return value
}

func (stack *MultiStack) PopTail() any {
	if stack.tail >= stack.max {
		return nil
	}
	
	value := stack.items[stack.tail]
	stack.items[stack.tail] = nil
	stack.tail++
	return value
}

func (stack *MultiStack) PeekHead() any {
	if stack.head <= -1 {
		return nil
	}
	
	return stack.items[stack.head]
}

func (stack *MultiStack) PeekTail() any {
	if stack.tail >= stack.max {
		return nil
	}
	
	return stack.items[stack.tail]
}
