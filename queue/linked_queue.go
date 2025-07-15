package queue

import "fmt"

type Node[T any] struct {
	data T
	next *Node[T]
}

type LinkedQueue[T any] struct {
	front *Node[T]
	rear *Node[T]
	size uint
}

func (q *LinkedQueue[T]) Enqueue(value T) {
	n := &Node[T]{data: value, next: nil}
	if q.rear != nil {
		q.rear.next = n
		q.rear = n
	} else {
		q.rear = n
		q.front = n
	}
	
	q.size++
}

func (q *LinkedQueue[T]) Dequwuw() (T, error) {
	if q.front == nil {
		return *new(T), fmt.Errorf("empty queue")
	}
	n := q.front.data
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil
	}
	
	q.size--
	return n, nil
}

func (q *LinkedQueue[T]) Peek() (T, error) {
	if q.front == nil {
		return *new(T), fmt.Errorf("empty queue")
	}
	
	return q.front.data, nil
}

func (q *LinkedQueue[T]) Size() uint {
	return q.size
}
