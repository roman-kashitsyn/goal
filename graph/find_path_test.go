package graph

import (
	"testing"
)

func verifyShortestPath(t *testing.T, exp, act []Vertex) {
	if len(exp) != len(act) {
		t.Fatalf("Expected to find path of len(%d), %v; got path of len(%d): %v",
			len(exp), exp, len(act), act)
	}
	for i, v := range act {
		if e := exp[i]; e != v {
			t.Fatalf("%d-th vertex expected to be %v, got %v", i, e, v)
		}
	}
}

func TestSimpleShortestPath(t *testing.T) {
	g := NewAdjacencyList(6, false)
	// 0 --> 1 --> 2 --> 3
	//       |           |
	//       |           V
	//       +---> 4 --> 5
	g.AddEdge(0, 1).AddEdge(1, 2).AddEdge(2, 3)
	g.AddEdge(1, 4).AddEdge(4, 5).AddEdge(3, 5)
	path, err := ShortestPath(g, 0, 5)
	if err != nil {
		t.Error(err)
	}
	expected := []Vertex{0, 1, 4, 5}
	verifyShortestPath(t, expected, path)
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

func TestSimpleDijkstraSearch(t *testing.T) {
	g := NewAdjacencyList(5, true)
	// 0 --> 1 --> 2 --> 3
	// |  1     1     1  ^
	// |                 |
	// +---> 4 ----------+
	//   1         3
	g.AddEdge(0, 1).AddEdge(1, 2).AddEdge(2, 3)
	g.AddEdge(0, 4).AddEdge(4, 3)
	weights := [][]int{
		[]int{0, 1, 0, 0, 1},
		[]int{0, 0, 1, 0, 0},
		[]int{0, 0, 0, 1, 0},
		[]int{0, 0, 0, 0, 0},
		[]int{0, 0, 0, 3, 0},
	}
	path, e := DijkstraShortestPath(g, 0, 3, WeightByTable(weights))
	if e != nil {
		t.Error(e)
	}
	expected := []Vertex{0, 1, 2, 3}
	verifyShortestPath(t, expected, path)
}
