package tree

import "cmp"

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
