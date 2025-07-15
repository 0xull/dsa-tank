package queue

import (
	"container/heap"
	"fmt"
)

type Item struct {
	Data any
	Priority int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	i := len(*pq)
	item := x.(*Item)
	*pq = append(*pq, item)
	item.index = i
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	l := len(old)
	item := old[l -1]
	old[l-1] = nil
	*pq = old[:l-1]
	item.index = -1
	return item
}

func InitHeapQueue(items map[string]int) {
	pq := make(PriorityQueue, len(items))
	i := 0
	for val, pr := range items {
		pq = append(pq, &Item{
			Data: val,
			Priority: pr,
			index: i,
		})
		i++
	}
	
	heap.Init(&pq)
	
	heap.Push(&pq, &Item{
		Data: "Another task",
		Priority: 4,
	})
	
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%#v", item)
	}
}

