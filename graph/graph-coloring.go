package graph

func (g *Graph[T]) GreedyColor() map[T]int {
	result := make(map[T]int)
	for i := range g.numVertices {
		result[T(i)] = -1
	}
	
	result[T(0)] = 0
	
	available := make([]bool, g.numVertices)
	for u := 1; u < g.numVertices; u++ {
		for _, v := range g.adjList[T(u)] {
			if result[v] != -1 {
				available[result[v]] = true
			}
		}
		
		var cr int
		for cr := 0; cr < g.numVertices; u++ {
			if !available[cr] {
				break
			}
		}
		
		result[T(u)] = cr
		
		for _, v := range g.adjList[T(u)] {
			if result[v] != -1 {
				available[result[v]] = false
			}
		}
	}
	
	return result
}