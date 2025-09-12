package graph

import (
	"cmp"
	"container/heap"
	"sort"
)

// Edge represents a connection between two vertices with specific weight.
type Edge[T cmp.Ordered] struct {
	From   T
	To     T
	Weight int
}

// WeightedGraph uses an adjacency list to represent a weighted, undirected graph.
type WeightedGraph[T cmp.Ordered] struct {
	numVertices int
	adjList     map[T][]Edge[T]
}

func NewWeightedGraph[T cmp.Ordered](numVertices int) *WeightedGraph[T] {
	return &WeightedGraph[T]{
		numVertices: numVertices,
		adjList:     make(map[T][]Edge[T]),
	}
}

func (g *WeightedGraph[T]) AddEdge(from T, to T, weight int) {
	g.adjList[from] = append(g.adjList[from], Edge[T]{From: from, To: to, Weight: weight})
	g.adjList[to] = append(g.adjList[to], Edge[T]{From: to, To: from, Weight: weight})
}

// EdgePriorityQueue implements heap.Interface for a slice of Edges. A Min-Heap based on
// edge weights.
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

// PrimsMST finds the Minimum Spanning Tree of a graph using the Prim's algorithm.
func (g *WeightedGraph[T]) PrimsMST(startVertex T) ([]Edge[T], int) {
	visited := make(map[T]bool)
	mst := []Edge[T]{}
	totalWeight := 0
	pq := make(EdgePriorityQueue[T], 0)

	addEdges := func(u T) {
		visited[u] = true
		for _, edge := range g.adjList[u] {
			if !visited[edge.To] {
				heap.Push(&pq, &Edge[T]{From: u, To: edge.To, Weight: edge.Weight})
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

		mst = append(mst, *edge)
		totalWeight += edge.Weight

		addEdges(v)
	}

	return mst, totalWeight
}

type DSU[T cmp.Ordered] struct {
	parent        []int
	size          []int
	vertexToIndex map[T]int
}

func NewDSU[T cmp.Ordered](vertices []T) *DSU[T] {
	count := len(vertices)
	parent := make([]int, count)
	size := make([]int, count)
	vertexToIndex := make(map[T]int, count)

	for i, v := range vertices {
		parent[i] = i
		size[i] = 1
		vertexToIndex[v] = i
	}

	return &DSU[T]{
		parent:        parent,
		size:          size,
		vertexToIndex: vertexToIndex,
	}
}

// findRec is the internal recursive helper tat works with integer indices.
func (dsu *DSU[T]) findRec(i int) int {
	if dsu.parent[i] != i {
		dsu.parent[i] = dsu.findRec(dsu.parent[i])
	}
	return dsu.parent[i]
}

// Find returns the representative of the set containing the vertex v.
func (dsu *DSU[T]) Find(v T) int {
	index, ok := dsu.vertexToIndex[v]
	if !ok {
		return -1
	}
	return dsu.findRec(index)
}

// Union merges the sets containing vertices u and v.
func (dsu *DSU[T]) Union(u, v T) {
	rootU := dsu.Find(u)
	rootV := dsu.Find(v)

	if rootU == rootV {
		return
	}

	// Union by size
	if dsu.size[rootU] < dsu.size[rootV] {
		dsu.parent[rootU] = rootV
		dsu.size[rootV] += dsu.size[rootU]
	} else {
		dsu.parent[rootV] = rootU
		dsu.size[rootU] += dsu.size[rootV]
	}
}

func (g *WeightedGraph[T]) KruskalsMST() ([]Edge[T], int) {
	var vertices []T
	seenVertices := make(map[T]bool)
	for v := range g.adjList {
		if !seenVertices[v] {
			vertices = append(vertices, v)
			seenVertices[v] = true
		}
	}

	var edges []Edge[T]
	for from, neighbour := range g.adjList {
		for _, edge := range neighbour {
			if from < edge.To {
				edges = append(edges, edge)
			}
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})
	
	mst := []Edge[T]{}
	totalWeight := 0
	dsu := NewDSU(vertices)
	
	for _, edge := range edges {
		if len(mst) == g.numVertices {
			break
		}
		
		if dsu.Find(edge.From) != dsu.Find(edge.To) {
			mst = append(mst, edge)
			totalWeight += edge.Weight
			dsu.Union(edge.From, edge.To)
		}
	}
	
	return mst, totalWeight
}
