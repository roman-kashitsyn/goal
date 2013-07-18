package graph

import "testing"

func verifyTopologicalSort(t *testing.T, x, y []Vertex) {
	if len(x) != len(y) {
		t.Fatalf("Expected slice of length %d, but length is %d", len(x), len(y))
	}
	for i, v := range x {
		if yv := y[i]; yv != v {
			t.Fatalf("Expected to find vertex %d on position %d, got vertex %d", v, i, yv)
		}
	}
}

func checkContainsLoop(t *testing.T, g Graph) {
	_, err := TopologicalSort(g)
	t.Logf("%v", err)
	if err == nil {
		t.Fatalf("Graph loop must be detected")
	}
}

func TestTopoSort(t *testing.T) {
	g := NewAdjacencyList(5, true)
	// 0 --> 1 --> 2
	// |     |
	// V     V
	// 3 --> 4
	g.AddEdge(0, 1).AddEdge(1, 2).AddEdge(1, 4)
	g.AddEdge(0, 3).AddEdge(3, 4)

	sorted, err := TopologicalSort(g)
	if err != nil {
		t.Error(err)
	}
	expected := []Vertex{2, 4, 1, 3, 0}
	verifyTopologicalSort(t, expected, sorted)
}

func TestCyclicGraphSort(t *testing.T) {
	g := NewAdjacencyList(5, true)
	// 0 --> 1 --> 2
	// |     ^     |
	// |     |     V
	// +---> 4 <-- 3
	g.AddEdge(0, 1).AddEdge(1, 2).AddEdge(2, 3)
	g.AddEdge(3, 4).AddEdge(4, 1).AddEdge(0, 4)
	checkContainsLoop(t, g)
}

func TestLongCyclicGraphSort(t *testing.T) {
	n := 8
	g := NewAdjacencyList(n, true)
	// 0 --> 1 --> 2 --> 3
	// ^                 |
	// |                 V
	// 7 <-- 6 <-- 5 <-- 4
	for i := 0; i < n; i++ {
		g.AddEdge(Vertex(i), Vertex((i+1)%n))
	}
	checkContainsLoop(t, g)
}
