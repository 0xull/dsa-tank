package stackarray

// StackArray is an array based implementation of the stack ADT
type StackArray struct {
	items []any
	top int
	max uint
}

// NewStackArray initializes and returns a pointer instance of StackArray
func NewStackArray(capacity uint) *StackArray {
	if capacity == 0 {
		capacity = 10
	}
	
	return &StackArray{
		items: make([]any, capacity),
		top: -1,
		max: capacity
	}
}


