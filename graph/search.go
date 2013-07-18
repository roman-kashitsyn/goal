package graph

import (
	"github.com/roman-kashitsyn/goal/queue"
)

func BreadthFirstSearch(g Graph, start Vertex, t Traverser) *Context {
	q := queue.NewLinkedQueue()
	n := g.NumVertices()
	c := makeContext(n)

	q.Enqueue(start)
	c.discovered[start] = true
	c.parents[start] = start
	directed := g.IsDirected()

	for !q.Empty() {
		v := q.Dequeue().(Vertex)
		t.OnEnter(c, v)

		for _, a := range g.AdjacentOf(v) {
			if !c.processed[a] || directed {
				t.OnEdge(c, v, a)
			}
			if !c.discovered[a] {
				q.Enqueue(a)
				c.discovered[a] = true
				c.parents[a] = v
			}
		}
		c.processed[v] = true
		t.OnExit(c, v)
	}

	t.OnFinish(c)
	return c
}

func DfsWithContext(g Graph, v Vertex, t Traverser, c *Context) *Context {
	if c == nil {
		c = makeContext(g.NumVertices())
	}

	c.discovered[v] = true
	t.OnEnter(c, v)

	for _, a := range g.AdjacentOf(v) {
		if !c.discovered[a] {
			c.parents[a] = v
			t.OnEdge(c, v, a)
			DfsWithContext(g, a, t, c)
		} else if !c.processed[a] || g.IsDirected() {
			t.OnEdge(c, v, a)
		}
	}

	c.processed[v] = true
	t.OnExit(c, v)
	return c
}

func DepthFirstSearch(g Graph, start Vertex, t Traverser) *Context {
	return DfsWithContext(g, start, t, nil)
}
