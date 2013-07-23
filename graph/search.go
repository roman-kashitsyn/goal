package graph

import (
	"github.com/roman-kashitsyn/goal/queue"
	"fmt"
)

type SearchFunc func(Graph, Vertex, Traverser, Context) Context

// Context interface represents traversal context. An instance of this
// interface is provided to any traverser during graph traversal.
type Context interface {
	ParentOf(v Vertex) Vertex
	SetParent(v, p Vertex)

	IsDiscovered(v Vertex) bool
	MarkDiscovered(v Vertex)

	IsProcessed(v Vertex) bool
	MarkProcessed(v Vertex)
}

// Traverser represents a type that reacts on events fired during
// graph traversal.
type Traverser interface {
	OnEnter(c Context, v Vertex)
	OnEdge(c Context, x, y Vertex)
	OnExit(c Context, v Vertex)
	OnFinish(c Context)
}

type SimpleContext struct {
	parents []Vertex
	discovered []bool
	processed []bool
}

func makeContext(numVertices int) *SimpleContext {
	parents := make([]Vertex, numVertices)
	discovered := make([]bool, numVertices)
	processed := make([]bool, numVertices)
	return &SimpleContext{parents, discovered, processed}
}

func (c *SimpleContext) ParentOf(v Vertex) Vertex {
	return c.parents[v]
}

func (c *SimpleContext) SetParent(v, p Vertex) {
	c.parents[int(v)] = p
}

func (c *SimpleContext) IsDiscovered(v Vertex) bool {
	return c.discovered[int(v)]
}

func (c *SimpleContext) MarkDiscovered(v Vertex) {
	c.discovered[int(v)] = true
}

func (c *SimpleContext) IsProcessed(v Vertex) bool {
	return c.processed[int(v)]
}

func (c *SimpleContext) MarkProcessed(v Vertex) {
	c.processed[int(v)] = true
}

type NoOpTraverser struct{}

func (t NoOpTraverser) OnEnter(c Context, v Vertex) {
}

func (t NoOpTraverser) OnEdge(c Context, x, y Vertex) {
}

func (t NoOpTraverser) OnExit(c Context, v Vertex) {
}

func (t NoOpTraverser) OnFinish(c Context) {
}

type PrintingTraverser struct{}

func (t PrintingTraverser) OnEnter(c Context, v Vertex) {
	fmt.Println("[+V] Entering vertex ", v)
}

func (t PrintingTraverser) OnEdge(c Context, x, y Vertex) {
	fmt.Printf("[+E] Entering edge %v -> %v\n", x, y)
}

func (t PrintingTraverser) OnExit(c Context, v Vertex) {
	fmt.Println("[-V] Exiting vertex ", v)
}

func (t PrintingTraverser) OnFinish(c Context) {
}

// BreadthFirstSearch runs breadth-first search on a given graph.
//
// Time complexity: O(|E| + |V|) where |E| is number of graph edges
// and |V| is number of graph vestices.
func BreadthFirstSearch(g Graph, start Vertex, t Traverser) Context {
	q := queue.NewLinkedQueue()
	n := g.NumVertices()
	c := makeContext(n)

	q.Enqueue(start)
	c.MarkDiscovered(start)
	c.SetParent(start, start)
	directed := g.IsDirected()

	for !q.Empty() {
		v := q.Dequeue().(Vertex)
		t.OnEnter(c, v)

		for _, a := range g.AdjacentOf(v) {
			if !c.IsProcessed(a) || directed {
				t.OnEdge(c, v, a)
			}
			if !c.discovered[a] {
				q.Enqueue(a)
				c.MarkDiscovered(a)
				c.SetParent(a, v)
			}
		}
		c.MarkProcessed(v)
		t.OnExit(c, v)
	}

	t.OnFinish(c)
	return c
}

// DfsWithContext runs depth-first search on a given graph using
// specified context. If context is nil, a new one will be created.
//
// Time complexity: O(|E| + |V|) where |E| is number of graph edges
// and |V| is number of graph vestices.
func DfsWithContext(g Graph, v Vertex, t Traverser, c Context) Context {
	if c == nil {
		c = makeContext(g.NumVertices())
	}

	c.MarkDiscovered(v)
	t.OnEnter(c, v)

	for _, a := range g.AdjacentOf(v) {
		if !c.IsDiscovered(a) {
			c.SetParent(a, v)
			t.OnEdge(c, v, a)
			DfsWithContext(g, a, t, c)
		} else if !c.IsProcessed(a) || g.IsDirected() {
			t.OnEdge(c, v, a)
		}
	}

	c.MarkProcessed(v)
	t.OnExit(c, v)
	return c
}

// DepthFirstSearch runs depth-first search on a given graph.
func DepthFirstSearch(g Graph, start Vertex, t Traverser) Context {
	return DfsWithContext(g, start, t, nil)
}

func ScanAllVertices(g Graph, t Traverser, f SearchFunc) Context {
	n := g.NumVertices()
	c := makeContext(n)
	for i := 0; i < n; i++ {
		v := Vertex(i)
		if !c.IsProcessed(v) {
			f(g, v, t, c)
		}
	}
	return c
}