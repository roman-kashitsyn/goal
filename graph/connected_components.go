package graph

type SccTraverser interface {
	OnComponentStart()
	OnComponentElement(v Vertex)
}

type ComponentRecorder struct {
	Components [][]Vertex
}

func NewComponentRecorder() *ComponentRecorder {
	return &ComponentRecorder{make([][]Vertex, 0)}
}

func (c *ComponentRecorder) OnComponentStart() {
	c.Components = append(c.Components, make([]Vertex, 0))
}

func (c *ComponentRecorder) OnComponentElement(v Vertex) {
	n := len(c.Components) - 1
	c.Components[n] = append(c.Components[n], v)
}

func TarjanScc(g Graph, sccTrav SccTraverser) {
	n := g.NumVertices()
	index := 0
	indexOf := make([]int, n)
	lowlinkOf := make([]int, n)
	stack := make([]Vertex, 0, n)

	for i, _ := range indexOf {
		indexOf[i] = -1
	}

	var connect func(Vertex)
	
	connect = func(v Vertex) {
		i := int(v)
		indexOf[i] = index
		lowlinkOf[i] = index
		index++
		sccStackPush(&stack, v)

		for _, a := range g.AdjacentOf(v) {
			if indexOf[int(a)] == -1 {
				connect(a)
				lowlinkOf[i] = intMin(lowlinkOf[i],lowlinkOf[int(a)])
			} else if sccStackContains(stack, a) {
				lowlinkOf[i] = intMin(lowlinkOf[i], indexOf[int(a)])
			}
		}

		if lowlinkOf[i] == indexOf[i] {
			sccTrav.OnComponentStart()
			done := false
			for !done {
				top := sccStackPop(&stack)
				sccTrav.OnComponentElement(top)
				done = (top == v)
			}
		}
	}

	for i := 0; i < n; i++ {
		if indexOf[i] == -1 {
			connect(Vertex(i))
		}
	}
}

func sccStackContains(stack []Vertex, v Vertex) bool {
	for _, s := range stack {
		if v == s {
			return true
		}
	}
	return false
}

func sccStackPush(stack *[]Vertex, v Vertex) {
	*stack = append(*stack, v)
}

func sccStackPop(stack *[]Vertex) (v Vertex) {
	n := len(*stack) - 1
	*stack, v = (*stack)[:n], (*stack)[n]
	return v
}

func intMin(x, y int) int {
	m := x
	if y < m {
		m = y
	}
	return m
}
