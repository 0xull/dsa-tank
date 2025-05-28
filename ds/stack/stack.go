package stackarray

import (
	"fmt"
)

// StackArray is an array based implementation of the stack ADT
type StackArray struct {
	items []any
	top int
	max int
}

// NewStackArray initializes and returns a pointer instance of StackArray
func NewStackArray(capacity int) *StackArray {
	if capacity <= 0 {
		capacity = 10
	}
	
	return &StackArray{
		items: make([]any, capacity),
		top: -1,
		max: capacity,
	}
}

// Push adds a new item to the top of the stack.
func (stack *StackArray) Push(value any) error {
	if stack.top >= (stack.max - 1) {
		return fmt.Errorf("stack overflow")
	}
	stack.top++
	stack.items[stack.top] = value
	return nil
}

// Pop removes and returns the item at the top of the stack.
func (stack *StackArray) Pop() (any, error) {
	if stack.top == -1 {
		return nil, fmt.Errorf("stack underflow")
	}
	item := stack.items[stack.top]
	stack.items[stack.top] = nil
	stack.top--
	return item, nil
}

// Peek returns the item at the top of the stack without removing it.
func (stack *StackArray) Peek() any {
	if stack.top == -1 {
		return nil
	}
	
	return stack.items[stack.top]
}
