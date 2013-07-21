// Package graph provides interfaces and data structures for
// representation of graphs and implementation of common graph
// algorithms.
package graph

type Vertex int

type Graph interface {
	// AdjacentOf returns all vertices adjacent to given vertex v.
	AdjacentOf(v Vertex) []Vertex

	// HasEdge checks whether an edge is present in graph.
	HasEdge(x, y Vertex) bool

	// Returns true if graph is a directed graph.
	IsDirected() bool

	// NumVertices returns number of vertices in a graph.
	NumVertices() int
}

type MutableGraph interface {
	Graph
	AddEdge(x, y Vertex) MutableGraph
}