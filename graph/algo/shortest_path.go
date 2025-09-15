package graph

import (
	"cmp"
	"container/heap"
	"fmt"
	"math"
)

type Item[T cmp.Ordered] struct {
	vertex   T
	priority int
	index    int
}

type PriorityQueue[T cmp.Ordered] []*Item[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*Item[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

func (g WeightedGraph[T]) Dijkstra(startVertex T) (map[T]int, map[T]T) {
	distance := make(map[T]int)
	predescessors := make(map[T]T)
	pq := make(PriorityQueue[T], 0)

	for i := range g.numVertices {
		distance[T(i)] = math.MaxInt32
	}
	distance[startVertex] = 0

	heap.Push(&pq, &Item[T]{vertex: startVertex, priority: 0})
	
	for pq.Len() > 0 {
		uitem := heap.Pop(&pq).(*Item[T])
		u := uitem.vertex
		
		for _, edge := range g.adjList[u] {
			v := edge.To
			weight := edge.Weight
			
			newDistance := distance[u] + weight
			if newDistance < distance[v] {
				distance[v] = newDistance
				predescessors[v] = u
				heap.Push(&pq, &Item[T]{vertex: v, priority: newDistance})
			}
		}
	}
	return distance, predescessors
}

func (g *WeightedGraph[T]) BellmanFord(startVertex T) (map[T]int, map[T]T, error) {
	distances := make(map[T]int)
	predescessors := make(map[T]T)
	
	for i := range g.numVertices {
		distances[T(i)] = math.MaxInt32
	}
	distances[startVertex] = 0
	
	var allEdges []Edge[T]
	for from, neighbors := range g.adjList {
		for _, edge := range neighbors {
			allEdges = append(allEdges, Edge[T]{From: from, To: edge.To, Weight: edge.Weight})
		}
	}
	
	for range g.numVertices-1 {
		for _, edge := range allEdges {
			u, v, weight := edge.From, edge.To, edge.Weight
			
			if distances[u] != math.MaxInt32 && distances[u]+weight < distances[v] {
				distances[v] = distances[u]+weight
				predescessors[v] = u
			}
		}
	}
	
	for _, edge := range allEdges {
		u, v, weight := edge.From, edge.To, edge.Weight
		if distances[u] != math.MaxInt32 && distances[u]+weight < distances[v] {
			return nil, nil, fmt.Errorf("graph contains a negative weight cycle")
		}
	}
	
	return distances, predescessors, nil
}
