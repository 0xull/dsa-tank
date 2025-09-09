package graph_test

import (
	"fmt"
	"testing"

	graph "github.com/IkehAkinyemi/dsa-tank/graph/algo"
)

func TestPrimsMST(t *testing.T) {
	g := graph.NewWeightGraph[int](4)
	g.AddEdge(0, 1, 2)
	g.AddEdge(0, 3, 4)
	g.AddEdge(0, 2, 5)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3, 3)
	g.AddEdge(2, 3, 6)
	
	mst, totalWeight := g.PrimsMST(0)
	fmt.Println("Edge in the Minimum Spanning Tree:")
	for _, edge := range mst {
		fmt.Printf("From: %d, To: %d, Weight: %d\n", edge.From, edge.To, edge.Weight)
	}
	fmt.Printf("Total Minimum weight: %d\n", totalWeight)
}
