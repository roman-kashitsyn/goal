package graph

type Vertex int

type Graph interface {
	AdjacentOf(v Vertex) []Vertex
	HasEdge(x, y Vertex) bool
	IsDirected() bool
	NumVertices() int
}

type Context struct {
	parents []Vertex
	discovered []bool
	processed []bool
}

type Traverser interface {
	OnEnter(c *Context, v Vertex) bool
	OnEdge(c *Context, x, y Vertex) bool
	OnExit(c *Context, v Vertex) bool
	OnFinish(c *Context)
}

type AdjacencyList struct {
	vertices [][]Vertex
	directed bool
}

func NewAdjacencyList(numVertices int, directed bool) *AdjacencyList {
	v := make([][]Vertex, numVertices)
	return &AdjacencyList{v, directed}
}

func (al *AdjacencyList) addSingleEdge(x, y Vertex) {
	edges := al.vertices[x]
	if edges == nil {
		edges = make([]Vertex, 0)
	}
	edges = append(edges, y)
	al.vertices[x] = edges
}

func (al *AdjacencyList) AddEdge(x, y Vertex) *AdjacencyList {
	al.addSingleEdge(x, y)
	if al.directed {
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

func makeContext(numVertices int) *Context {
	parents := make([]Vertex, numVertices)
	discovered := make([]bool, numVertices)
	processed := make([]bool, numVertices)
	return &Context{parents, discovered, processed}
}

func (c *Context) ParentOf(v Vertex) Vertex {
	return c.parents[v]
}

func (c *Context) IsDiscovered(v Vertex) bool {
	return c.discovered[int(v)]
}

func (c *Context) IsProcessed(v Vertex) bool {
	return c.processed[int(v)]
}

