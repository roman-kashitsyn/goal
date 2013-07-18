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
		if !t.OnEnter(c, v) {
			return c
		}
		for _, a := range g.AdjacentOf(v) {
			if !c.processed[a] || directed {
				if !t.OnEdge(c, v, a) {
					return c
				}
			}
			if !c.discovered[a] {
				q.Enqueue(a)
				c.discovered[a] = true
				c.parents[a] = v
			}
		}
		c.processed[v] = true
		if !t.OnExit(c, v) {
			return c
		}
	}

	t.OnFinish(c)
	return c
}

func dfs(g Graph, v Vertex, t Traverser, c *Context) {
	c.discovered[v] = true
	if !t.OnEnter(c, v) {
		return
	}

	for _, a := range g.AdjacentOf(v) {
		if !c.discovered[a] {
			c.parents[a] = v
			if !t.OnEdge(c, v, a) {
				return
			}
			dfs(g, a, t, c)
		} else if !c.processed[a] || g.IsDirected() {
			if !t.OnEdge(c, v, a) {
				return
			}
		}
	}

	c.processed[v] = true
	if !t.OnExit(c, v) {
		return
	}
}

func DepthFirstSearch(g Graph, start Vertex, t Traverser) *Context {
	n := g.NumVertices()
	c := makeContext(n)
	dfs(g, start, t, c)
	return c
}
