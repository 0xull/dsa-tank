package graph

import "fmt"

func (g *Graph[T]) TopologicalSort() ([]T, error) {
	inDegree := make(map[T]int)
	for i := range g.numVertices {
		inDegree[T(i)] = 0
	}
	
	for _, neighbors := range g.adjList {
		for _, v := range neighbors {
			inDegree[v]++
		}
	}
	
	queue := []T{}
	for i := range g.numVertices {
		if inDegree[T(i)] == 0 {
			queue = append(queue, T(i))
		}
	}
	
	sortedOrder := []T{}
	visitedCount := 0
	
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		sortedOrder = append(sortedOrder, u)
		visitedCount++
		
		for _, v := range g.adjList[u] {
			inDegree[v]--
			if inDegree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	
	if visitedCount != g.numVertices {
		return nil, fmt.Errorf("graph contains a cycle, topological sort not possible")
	}
	
	return sortedOrder, nil
}