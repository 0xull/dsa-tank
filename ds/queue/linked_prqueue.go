package queue

import "fmt"

type PRNode[T any] struct {
	data     T
	priority int
	next     *PRNode[T]
}

type LinkedPRQueue[T any] struct {
	head *PRNode[T]
	size int
}

func (q *LinkedPRQueue[T]) Enqueue(val T, priority int) {
	n := &PRNode[T]{val, priority, nil}
	if q.head == nil || q.head.priority > priority {
		n.next = q.head
		q.head = n
		q.size++
		return
	}
	curr := q.head
	for curr.next != nil && curr.next.priority <= priority {
		curr = curr.next
	}
	n.next = curr.next
	curr.next = n
	q.size++
}

func (q *LinkedPRQueue[T]) Dequeue() (T, error) {
	if q.head == nil {
		return *new(T), fmt.Errorf("empty queue")
	}
	v := q.head.data
	q.head = q.head.next
	q.size--
	return v, nil
}

func (q *LinkedPRQueue[T]) Peek() (T, error) {
	if q.head == nil {
		return *new(T), fmt.Errorf("empty queue")
	}
	return q.head.data, nil
}

func (q *LinkedPRQueue[T]) Size() int {
	return q.size
}
