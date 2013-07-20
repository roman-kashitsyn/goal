package graph

import (
	"testing"
)

func TestSimpleShortestPath(t *testing.T) {
	g := NewAdjacencyList(6, false)
	// 0 --> 1 --> 2 --> 3
	//       \           |
	//       \           V
	//       \---> 4 --> 5
	g.AddEdge(0, 1).AddEdge(1, 2).AddEdge(2, 3)
	g.AddEdge(1, 4).AddEdge(4, 5).AddEdge(3, 5)
	path, err := ShortestPath(g, 0, 5)
	if err != nil {
		t.Error(err)
	}
	expected := []Vertex{0, 1, 4, 5}
	for i, v := range path {
		if e := expected[i]; e != v {
			t.Fatalf("%d-th vertex expected to be %v, got %v", i, e, v)
		}
	}
}

func TestUnconnectedPathSearch(t *testing.T) {
	g := NewAdjacencyList(3, true)
	// 0 --> 1 <-- 2
	g.AddEdge(0, 1).AddEdge(2, 1)
	p, err := ShortestPath(g, 0, 2)
	if err == nil {
		t.Fatalf("ShortestPath: no path exists, but error is nil, path: %v", p)
	}
}