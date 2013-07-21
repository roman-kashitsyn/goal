package graph

type AdjacencyMatrix struct {
	matrix []bool
	numVertices int
	directed bool
}

func NewAdjacencyMatrix(numVertices int, directed bool) *AdjacencyMatrix {
	size := numVertices * numVertices
	matrix := make([]bool, size)
	return &AdjacencyMatrix{matrix, numVertices, directed}
}

func (m *AdjacencyMatrix) AddEdge(x, y Vertex) MutableGraph {
	m.matrix[m.edgeIndex(x, y)] = true
	if m.directed {
		m.matrix[m.edgeIndex(y, x)] = true
	}
	return m
}

func (m *AdjacencyMatrix) AdjacentOf(v Vertex) []Vertex {
	start := int(v) * m.numVertices
	end := start + m.numVertices
	row := m.matrix[start:end]
	vertices := make([]Vertex, 0, m.numVertices / 2)
	for i, present := range row {
		if present {
			vertices = append(vertices, Vertex(i))
		}
	}
	return vertices
}

func (m *AdjacencyMatrix) HasEdge(x, y Vertex) bool {
	return m.matrix[m.edgeIndex(x, y)]
}

func (m *AdjacencyMatrix) IsDirected() bool {
	return m.directed
}

func (m *AdjacencyMatrix) NumVertices() int {
	return m.numVertices
}

func (m *AdjacencyMatrix) edgeIndex(x, y Vertex) int {
	return int(x) * m.numVertices + int(y)
}