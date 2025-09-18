package graph

// Kosaraju finds all Strongly Connected Graphs in a directed graph.
func (g *Graph[T]) Kosaraju() [][]T {
	stack := []T{}
	visited := make(map[T]bool)
	for i := range g.numVertices {
		if !visited[T(i)] {
			g.fillOrder(T(i), visited, &stack)
		}
	}

	reversedGraph := g.getTranspose()

	var allSCC [][]T
	visited = make(map[T]bool)
	for len(stack) > 0 {
		u := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		if !visited[u] {
			var singleSCC []T
			reversedGraph.dfsForSCC(u, visited, &singleSCC)
			allSCC = append(allSCC, singleSCC)
		}
	}
	
	return allSCC
}

// fillOrder is the first DFS in Kosaraju algorithm that populates the stack.
func (g *Graph[T]) fillOrder(u T, visited map[T]bool, stack *[]T) {
	visited[u] = true
	for _, v := range g.adjList[u] {
		if !visited[v] {
			g.fillOrder(v, visited, stack)
		}
	}
	*stack = append(*stack, u)
}

// getTranspose returns a new graph with all the edge directions reversed.
func (g *Graph[T]) getTranspose() *Graph[T] {
	reversed := NewGraph[T](g.numVertices)
	for u, neighbors := range g.adjList {
		for _, v := range neighbors {
			reversed.AddDirectedEdge(v, u)
		}
	}
	return reversed
}

// dfsForSCC is the second DFS in Kosaraju, which collects nodes for one SCC.
func (g *Graph[T]) dfsForSCC(u T, visited map[T]bool, singleSCC *[]T) {
	visited[u] = true
	*singleSCC = append(*singleSCC, u)
	for _, v := range g.adjList[u] {
		if !visited[v] {
			g.dfsForSCC(v, visited, singleSCC)
		}
	}
}
