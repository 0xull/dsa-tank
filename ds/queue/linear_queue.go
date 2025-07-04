package queue

import "fmt"

type ArrayQueue[T any] struct {
	items []T
	front int
	rear int
	capacity int
}

func NewArrayQueue[T any](capacity int) *ArrayQueue[T] {
	if capacity <= 0 {
		capacity = 10
	}
	
	return &ArrayQueue[T]{
		items: make([]T, capacity),
		front: -1,
		rear: -1,
		capacity: capacity,
	}
}

func (q *ArrayQueue[T]) Enqueue(value T) error { 
	if q.rear == q.capacity-1 {
		return fmt.Errorf("queue overflow")
	}
	
	q.rear++
	q.items[q.rear] = value
	
	if q.front == -1 {
		q.front = 0
	}

	return nil 
}

func (q *ArrayQueue[T]) Dequeue() (T, error) {
	// we check for underflow in two situation:
	//   1. Initial empty set when front == -1
	//   2. When front > rear, meaning last item is dequeued.
	//      But we're immediately checking for that and resetting values as necessary
	//      meaning such check here would never occur, as a result useless.
	if q.front == -1 {
		return *new(T), fmt.Errorf("queue underflow")
	}
	
	v := q.items[q.front]
	q.items[q.front] = *new(T)
	q.front++
	
	if q.front > q.rear {
		q.front = -1
		q.rear = -1
	}
	
	return v, nil 
}

func (q *ArrayQueue[T]) Peek() (T, error) { 
	if 	q.front == -1 {
		return *new(T), fmt.Errorf("empty queue")
	}
	
	return q.items[q.front], nil 
}

func (q *ArrayQueue[T]) Size() int {
	if q.front == -1 {
		return 0
	}
	return (q.rear - q.front) + 1
}