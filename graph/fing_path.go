package graph

import (
	"fmt"
	"math"
)

type WeightFunc func(x, y Vertex) int

type NotReachableError struct {
	start, end Vertex
}

func (e *NotReachableError) Error() string {
	return fmt.Sprintf("Vertex %v is not reachable from vertex %v", e.end, e.start)
}

func buildPath(parents []Vertex, start, end Vertex) []Vertex {
	n := 0
	for v := end; v != start; v = parents[v] {
		n++
	}
	path := make([]Vertex, n+1)
	for v := end; v != start; v = parents[v] {
		path[n] = v
		n--
	}
	return path
}

// ShortestPath find shortest path inp a graph assuming weights of
// all vestices are equal to 1.
func ShortestPath(g Graph, start, end Vertex) (path []Vertex, err error) {
	c := BreadthFirstSearch(g, start, NoOpTraverser{})
	if !c.IsDiscovered(end) {
		return nil, &NotReachableError{start, end}
	}
	return buildPath(c.parents, start, end), nil
}

// DijkstraShortestPath implements Dijkstra's shortest path algorithm.
//
// Time complexity: O(|E| + |V|^2) where |V| is number of graph vertices
func DijkstraShortestPath(g Graph, f, t Vertex, w WeightFunc) (path []Vertex, err error) {
	n := g.NumVertices()
	visited := make([]bool, n)
	distance := make([]int, n)
	parents := make([]Vertex, n)
	for i := 0; i < n; i++ {
		distance[i] = math.MaxInt32
		parents[i] = Vertex(-1)
	}
	distance[f] = 0
	v := f
	for !visited[v] {
		visited[v] = true

		for _, a := range g.AdjacentOf(v) {
			weight := w(v, a)
			if distance[a] < distance[v]+weight {
				distance[a] = distance[v] + weight
				parents[a] = v
			}
		}

		// Searching for vertex with minimal distance
		d := math.MaxInt32
		for i := 0; i < n; i++ {
			if !visited[i] && distance[i] < d {
				d = distance[i]
				v = Vertex(i)
			}
		}
	}

	if distance[t] == math.MaxInt32 {
		return nil, &NotReachableError{f, t}
	}
	return buildPath(parents, f, t), nil
}
