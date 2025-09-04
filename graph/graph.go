package graph

import (
	"cmp"
	"fmt"
)

type Graph[T cmp.Ordered] struct {
	numVertices int
	adjList     map[T][]T
}

func (g *Graph[T]) BFS(startVertex T) {
	visited := make(map[T]bool)
	queue := []T{}

	visited[startVertex] = true
	queue = append(queue, startVertex)

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		// process vertex
		fmt.Printf("%v\n", u)

		for _, v := range g.adjList[u] {
			if !visited[v] {
				queue = append(queue, v)
				visited[startVertex] = true
			}
		}
	}
}

func (g *Graph[T]) DFS(startVertex T) {
	visited := make(map[T]bool)
	g.dfsRec(startVertex, visited)
}

func (g *Graph[T]) dfsRec(u T, visited map[T]bool) {
	visited[u] = true

	// process vertex
	fmt.Printf("%v\n", u)

	for _, v := range g.adjList[u] {
		if !visited[v] {
			g.dfsRec(v, visited)
		}
	}
}

func (g *Graph[T]) DFSIter(startVertex T) {
	visited := make(map[T]bool)
	stack := []T{startVertex}

	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		if !visited[u] {			
			// process vertex
			fmt.Printf("%v\n", u)
			visited[u] = true
			
			for _, v := range g.adjList[u] {
				if !visited[v] {
					stack = append(stack, v)
				}
			}
		}
	}
}
