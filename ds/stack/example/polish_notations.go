package stack

type StackArray struct {
	items []rune
	top int
	capacity int
}

func NewStackArray(capacity int) *StackArray {
	if capacity <= 0 {
		capacity = 0
	}
	
	return &StackArray{
		items: make([]rune, capacity),
		top: -1,
		capacity: capacity,
	}
}

func (stack *StackArray) Push(value rune) {
	if stack.top == (stack.capacity -1) {
		return
	}
	
	stack.top++
	stack.items[stack.top] = value
}

func (stack *StackArray) Pop() rune {
	if stack.top == -1 {
		return 0
	}
	value := stack.items[stack.top]
	stack.items[stack.top] = 0
	stack.top--
	return value
}

