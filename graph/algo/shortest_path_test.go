package graph_test

import (
	"fmt"
	"math"
	"testing"

	graph "github.com/IkehAkinyemi/dsa-tank/graph/algo"
)

func TestDijskstra(t *testing.T) {
	// Create the graph from our example (A=0, B=1, C=2, D=3, E=4)
	g := graph.NewWeightedGraph[int](5)
	g.AddEdge(0, 1, 1)  // A -> B (1)
	g.AddEdge(0, 2, 4)  // A -> C (4)
	g.AddEdge(1, 4, 2)  // B -> E (2)
	g.AddEdge(1, 3, 9)  // B -> D (9)
	g.AddEdge(2, 3, 5)  // C -> D (5)

	startNode := 0
	distances, predecessors := g.Dijkstra(startNode)

	fmt.Printf("Shortest distances from vertex %d:\n", startNode)
	for i := range 5 {
		if distances[i] == math.MaxInt32 {
			fmt.Printf("  to vertex %d: infinity\n", i)
		} else {
			fmt.Printf("  to vertex %d: %d\n", i, distances[i])
		}
	}

	fmt.Printf("\nShortest path to vertex %d:\n", 3)
	path := []int{}
	curr := 3
	// Backtrack from the destination until we reach the start
	for curr != startNode {
		path = append([]int{curr}, path...)
		curr = predecessors[curr]
	}
	path = append([]int{startNode}, path...)
	fmt.Println(path)
}