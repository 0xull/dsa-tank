package queue

import "fmt"

type Deque[T any] struct {
	items []T
	left int
	right int
	capacity int
	size int
}

func NewDeque[T any](capacity int) *Deque[T] {
	if capacity <= 0 {
		capacity = 10
	}
	
	return &Deque[T]{
		items: make([]T, capacity),
		left: -1,
		right: -1,
		capacity: capacity,
		size: 0,
	}
}

func (q *Deque[T]) IsFull() bool {
	return q.size == q.capacity
}

func (q *Deque[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Deque[T]) Size() int {
	return q.size
}

func (q *Deque[T]) AddFront(val T) error {
	if q.IsFull() {
		return fmt.Errorf("queue overflow")
	}
	
	if q.IsEmpty() {
		q.left = 0
		q.right = 0
	} else {
		q.left = (q.left - 1 + q.capacity) % q.capacity
	}
	
	q.items[q.left] = val
	q.size++
	return nil
}

func (q *Deque[T]) AddRear(val T) error {
	if q.IsFull() {
		return fmt.Errorf("queue overflow")
	}
	
	if q.IsEmpty() {
		q.left = 0
		q.right = 0
	} else {
		q.right = (q.right + 1) % q.capacity
	}
	
	q.items[q.right] = val
	q.size++
	return nil
}

func (q *Deque[T]) RemoveFront() (T, error) {
	if q.IsEmpty() {
		return *new(T), fmt.Errorf("queue empty")
	}
	
	v := q.items[q.left]
	q.items[q.left] = *new(T)
	q.size--
	
	if q.size == 1 {
		q.left = -1
		q.right = -1
	} else {
		q.left = (q.left + 1) % q.capacity
	}
	
	return v, nil
}

func (q *Deque[T]) RemoveRear() (T, error) {
	if q.IsEmpty() {
		return *new(T), fmt.Errorf("queue empty")
	}
	
	v := q.items[q.right]
	q.items[q.right] = *new(T)
	q.size--
	
	if q.IsEmpty() {
		q.left = -1
		q.right = -1
	} else {
		q.right = (q.right - 1 + q.capacity) % q.capacity
	}
	
	return v, nil
}

func (q *Deque[T]) PeekFront() (T, error) {
	if q.IsEmpty() {
		return *new(T), fmt.Errorf("empty queue")
	}
	
	return q.items[q.left], nil
}

func (q *Deque[T]) PeekRear() (T, error) {
	if q.IsEmpty() {
		return *new(T), fmt.Errorf("empty queue")
	}
	
	return q.items[q.right], nil
}
