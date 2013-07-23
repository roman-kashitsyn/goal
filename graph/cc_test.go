package graph

import (
	"testing"
)


func TestSimpleConnectedComponents(t *testing.T) {
	g := NewAdjacencyList(8, true)
	// 0 --> 1  --> 2 <--> 3
	// ^    /       |      ^
	// |   /        |      |
	// |  /         V      V
	// 4 --> 5 <--> 6  <-- 7
	g.AddEdge(0, 1)
	g.AddEdge(1, 2).AddEdge(1, 4)
	g.AddEdge(2, 3).AddEdge(2, 6)
	g.AddEdge(3, 2).AddEdge(3, 7)
	g.AddEdge(4, 0).AddEdge(4, 5)
	g.AddEdge(5, 6)
	g.AddEdge(6, 5)
	g.AddEdge(7, 3).AddEdge(7, 6)

	trav := NewComponentRecorder()
	TarjanScc(g, trav)
	
	expected := [][]Vertex {
		[]Vertex{5, 6},
		[]Vertex{2, 3, 7},
		[]Vertex{0, 1, 4},
	}

	if len(trav.Components) != len(expected) {
		t.Logf("%v", trav.Components)
		t.Fatalf("Expected to discover %d components, found %d",
			len(expected), len(trav.Components))
	}
	for ci, c := range expected {
		comp := trav.Components[ci]
		for _, v := range c {
			if !sccStackContains(comp, v) {
				t.Fatalf("Expected to find v=%v in comp=%d(%v), got %v", v, ci, comp)
			}
		}
	}
	
}