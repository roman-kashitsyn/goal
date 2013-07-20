package graph

import (
	"fmt"
	"strings"
)

type topSort struct {
	out []Vertex
}

func buildLoopErrorMessage(c *Context, x, y Vertex) string {
	loop := make([]string, 0)

	loop = append(loop, fmt.Sprint(y))
	for v := x; v != y; v = c.ParentOf(v) {
		loop = append(loop, fmt.Sprint(v))
	}
	loop = append(loop, fmt.Sprint(y))

	n := len(loop)
	for i := 0; i < n/2; i++ {
		j := n - i - 1
		loop[i], loop[j] = loop[j], loop[i]
	}
	s := strings.Join(loop, " -> ")
	return fmt.Sprintf("Graph contains a cycle: %s", s)
}

func (t *topSort) OnEnter(c *Context, v Vertex) {
}

func (t *topSort) OnEdge(c *Context, x, y Vertex) {
	if c.IsDiscovered(y) && !c.IsProcessed(y) {
		panic(buildLoopErrorMessage(c, x, y))
	}
}

func (t *topSort) OnExit(c *Context, v Vertex) {
	t.out = append(t.out, v)
}

func (t *topSort) OnFinish(c *Context) {}

// TopologicalSort accepts direct acyclic graph and returns linear
// ordering of vertices such that for every directed edge uv from
// vertex u to vertex v, u comes before v in the ordering.
func TopologicalSort(dag Graph) (sorted []Vertex, e error) {
	defer func() {
		if err := recover(); err != nil {
			sorted = nil
			e = fmt.Errorf("%v", err)
		}
	}()

	n := dag.NumVertices()
	c := makeContext(n)
	t := &topSort{out: make([]Vertex, 0, n)}
	for i := 0; i < n; i++ {
		v := Vertex(i)
		if !c.IsProcessed(v) {
			DfsWithContext(dag, v, t, c)
		}
	}
	return t.out, nil
}
