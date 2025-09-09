package graph

import (
	"cmp"
	"container/heap"
)

// Edge represents a connection between two vertices with specific weight.
type Edge[T cmp.Ordered] struct {
	To     T
	Weight int
}

type WeightedGraph[T cmp.Ordered] struct {
	numVertices int
	adjList     map[T][]Edge[T]
}

type EdgePriorityQueue[T cmp.Ordered] []*Edge[T]

func (pq EdgePriorityQueue[T]) Len() int           { return len(pq) }
func (pq EdgePriorityQueue[T]) Less(i, j int) bool { return pq[i].Weight < pq[j].Weight }
func (pq EdgePriorityQueue[T]) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *EdgePriorityQueue[T]) Push(x any) {
	item := x.(*Edge[T])
	*pq = append(*pq, item)
}

func (pq *EdgePriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func (g *WeightedGraph[T]) PrimsMST(startVertex T) ([]Edge[T], int) {
	visited := make(map[T]bool)
	mst := []Edge[T]{}
	totalWeight := 0
	pq := make(EdgePriorityQueue[T], 0)

	addEdges := func(u T) {
		visited[u] = true
		for _, edge := range g.adjList[u] {
			if !visited[edge.To] {
				heap.Push(&pq, &Edge[T]{To: edge.To, Weight: edge.Weight})
			}
		}
	}
	
	addEdges(startVertex)
	
	for pq.Len() > 0 && len(mst) < g.numVertices-1 {
		edge := heap.Pop(&pq).(*Edge[T])
		v := edge.To
		
		if visited[v] {
			continue
		}
		
		mst := append(mst, *edge)
	}
}
