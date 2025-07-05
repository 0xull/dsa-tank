package queue

import "fmt"

type CircularQueue[T any] struct {
	items []T
	front int
	rear int
	capacity int
	size int
}

func NewCircularQueue[T any](capacity int) *CircularQueue[T] {
	if capacity <= 0 {
		capacity = 10
	}
	
	return &CircularQueue[T]{
		items: make([]T, capacity),
		front: -1,
		rear: -1,
		capacity: capacity,
		size: 0,
	}
}

func (q *CircularQueue[T]) IsFull() bool {
	return q.size == q.capacity
}

func (q *CircularQueue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *CircularQueue[T]) Size() int {
	return q.size
}

func (q *CircularQueue[T]) Enqueue(val T) error {
	if q.IsFull() {
		return fmt.Errorf("queue overflow")
	}
	
	if q.IsEmpty() {
		q.rear = 0
		q.front = 0
	} else {
		q.rear = (q.rear+1) % q.capacity
	}
	
	q.items[q.rear] = val
	q.size++
	
	return nil
}

func (q *CircularQueue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		return *new(T), fmt.Errorf("empty queue")
	}
	
	v := q.items[q.front]
	q.items[q.front] = *new(T)
	
	if q.size == 1 {
		q.front = -1
		q.rear = -1
	} else {
		q.front = (q.front+1) % q.capacity
	}
	
	q.size--
	return v, nil
}

func (q *CircularQueue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		return *new(T), fmt.Errorf("empty queue")
	}
	
	return q.items[q.front], nil
}
