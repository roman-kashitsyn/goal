package graph

// AdjacencyList is a simple implementation of adjacency list graph
// data structure.
type AdjacencyList struct {
	vertices [][]Vertex
	directed bool
}

func NewAdjacencyList(numVertices int, directed bool) *AdjacencyList {
	v := make([][]Vertex, numVertices)
	return &AdjacencyList{v, directed}
}

func (al *AdjacencyList) addSingleEdge(x, y Vertex) {
	al.vertices[x] = append(al.vertices[x], y)
}

func (al *AdjacencyList) AddEdge(x, y Vertex) MutableGraph {
	al.addSingleEdge(x, y)
	if !al.directed {
		al.addSingleEdge(y, x)
	}
	return al
}

func (al *AdjacencyList) AdjacentOf(v Vertex) []Vertex {
	return al.vertices[v]
}

func (al *AdjacencyList) HasEdge(x, y Vertex) bool {
	for _, v := range al.vertices[x] {
		if v == y {
			return true
		}
	}
	return false
}

func (al *AdjacencyList) IsDirected() bool {
	return al.directed
}

func (al *AdjacencyList) NumVertices() int {
	return len(al.vertices)
}
