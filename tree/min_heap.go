package tree

import (
	"cmp"
	"container/heap"
)

/* This sample would utilize the Go std library to implement a MinHeap. */

type MinHeap[T cmp.Ordered] []T

func (mnhp MinHeap[T]) Len() int { return len(mnhp) }

func (mnhp MinHeap[T]) Less(i, j int) bool { return mnhp[i] < mnhp[j] }

func (mnhp MinHeap[T]) Swap(i, j int) { mnhp[i], mnhp[j] = mnhp[j], mnhp[i] }

func (mnhp *MinHeap[T]) Push(data any) {
	i := data.(T)
	*mnhp = append(*mnhp, i)
}

func (mnhp *MinHeap[T]) Pop() any {
	old := *mnhp
	l := len(old)
	data := old[l-1]
	*mnhp = old[0 : l-1]
	return data
}

func BuildMinHeapFromSlice[T cmp.Ordered](slice []T) *MinHeap[T] {
	mnhp := MinHeap[T](slice)
	heap.Init(&mnhp)
	
	return  &mnhp
}
