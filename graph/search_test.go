package graph

import (
	"testing"
)

type recordingTraverser struct {
	order   []Vertex
	parents []Vertex
	test    *testing.T
}

func newTraverser(n int, t *testing.T) *recordingTraverser {
	o := make([]Vertex, 0, n)
	p := make([]Vertex, n)
	return &recordingTraverser{o, p, t}
}

func (t *recordingTraverser) OnEnter(c Context, v Vertex) {
	t.order = append(t.order, v)
	if !c.IsDiscovered(v) {
		t.test.Fatalf("Expected vertex %d to be discovered on enter", v)
	}
}

func (t *recordingTraverser) OnEdge(c Context, x, y Vertex) {
}

func (t *recordingTraverser) OnExit(c Context, v Vertex) {
	if !c.IsDiscovered(v) {
		t.test.Fatalf("Expected vertex %d to be discovered on exit", v)
	}
	if !c.IsProcessed(v) {
		t.test.Fatalf("Expected vertex %d to be processed on exit", v)
	}
	t.parents[v] = c.ParentOf(v)
}

func (t *recordingTraverser) OnFinish(c Context) {}

func verifySearch(t *testing.T, ord, par []Vertex, trav *recordingTraverser) {
	for i, v := range ord {
		if e := trav.order[i]; v != e {
			t.Fatalf("Expected %d-th vertex to be %d, got %d", i, v, e)
		}
	}
	for i, v := range par {
		if p := trav.parents[i]; v != p {
			t.Fatalf("Expected %d-th vertex parent to be %d, got %d", i, v, p)
		}
	}
}

func TestBreadthFirstSearch(t *testing.T) {
	// 0 -- 1 -- 2
	// |    |
	// +--- 3 -- 4
	test := func(g MutableGraph) {
		g.AddEdge(0, 1).AddEdge(1, 2).AddEdge(1, 3)
		g.AddEdge(0, 3).AddEdge(3, 4)
		trav := newTraverser(g.NumVertices(), t)
		BreadthFirstSearch(g, 0, trav)
		expectedOrder := []Vertex{0, 1, 3, 2, 4}
		expectedParents := []Vertex{0, 0, 1, 0, 3}
		verifySearch(t, expectedOrder, expectedParents, trav)
	}
	test(NewAdjacencyList(5, false))
	test(NewAdjacencyMatrix(5, false))
}

func TestDepthFirstSearch(t *testing.T) {
	// 0 -- 2 -- 4 --- 1
	// |         |     |
	// +--- 5 -- 3 ----+
	test := func(g MutableGraph) {
		g.AddEdge(0, 2).AddEdge(2, 4).AddEdge(4, 1).AddEdge(4, 3)
		g.AddEdge(0, 5).AddEdge(1, 3).AddEdge(3, 5)
		trav := newTraverser(g.NumVertices(), t)
		DepthFirstSearch(g, 0, trav)
		expectedOrder := []Vertex{0, 2, 4, 1, 3, 5}
		expectedParents := []Vertex{0, 4, 0, 1, 2, 3}
		verifySearch(t, expectedOrder, expectedParents, trav)
	}
	test(NewAdjacencyList(6, false))
	test(NewAdjacencyMatrix(6, false))
}
