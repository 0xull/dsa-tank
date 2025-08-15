package tree

import (
	"cmp"
	"errors"
)

type MaxHeap[T cmp.Ordered] struct {
	slice []T
}

func (mxhp *MaxHeap[T]) parent(index int) int {
	return (index - 1) / 2
}

func (mxhp *MaxHeap[T]) left(index int) int {
	return (index * 2) + 1
}

func (mxhp *MaxHeap[T]) right(index int) int {
	return (index * 2) + 2
}

func (mxhp *MaxHeap[T]) swap(i, j int) {
	mxhp.slice[i], mxhp.slice[j] = mxhp.slice[j], mxhp.slice[i]
}

func (mxhp *MaxHeap[T]) siftUp(index int) {
	for index > 0 && mxhp.slice[mxhp.parent(index)] < mxhp.slice[index] {
		mxhp.swap(mxhp.parent(index), index)
		index = mxhp.parent(index)
	}
}

func (mxhp *MaxHeap[T]) siftDown(index int) {
	maxIndent := index
	for {
		left := mxhp.left(index)
		right := mxhp.right(index)

		if left < len(mxhp.slice) && mxhp.slice[left] > mxhp.slice[maxIndent] {
			maxIndent = left
		}

		if right < len(mxhp.slice) && mxhp.slice[right] > mxhp.slice[maxIndent] {
			maxIndent = right
		}

		if index == maxIndent {
			break
		}

		mxhp.swap(index, maxIndent)
		index = maxIndent
	}
}

func (mxhp *MaxHeap[T]) Insert(data T) {
	mxhp.slice = append(mxhp.slice, data)
	mxhp.siftUp(len(mxhp.slice) - 1)
}

func (mxhp *MaxHeap[T]) Delete() (T, error) {
	if len(mxhp.slice) == 0 {
		return *new(T), errors.New("heap is empty")
	}

	data := mxhp.slice[0]
	lastIndex := len(mxhp.slice) - 1

	mxhp.slice[0] = mxhp.slice[lastIndex]
	mxhp.slice = mxhp.slice[:lastIndex]

	mxhp.siftDown(0)

	return data, nil
}

// Floyd algorithm build heap in O(n) complexity time.
func BuildMaxHeapFromSlice[T cmp.Ordered](slice []T) *MaxHeap[T] {
	heap := &MaxHeap[T]{slice: slice}

	lastNonLeafIndex := (len(heap.slice) / 2) - 1
	for i := lastNonLeafIndex; i <= 0; i-- {
		heap.siftDown(i)
	}
	
	return heap
}
