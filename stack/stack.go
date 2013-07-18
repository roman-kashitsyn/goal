package stack

type Interface interface {
	Push(e interface{})
	Pop() interface{}
	Top() interface{}
	Empty() bool
}

type SliceStack []interface{}

func NewSliceStack() *SliceStack {
	return &SliceStack{}
}

func NewSliceStackOfCapacity(c int) *SliceStack {
	s := new(SliceStack)
	*s = make([]interface{}, 0, c)
	return s
}

func (s *SliceStack) Push(e interface{}) {
	*s = append(*s, e)
}

func (s *SliceStack) Pop() (e interface{}) {
	i := len(*s) - 1
	*s, e = (*s)[:i], (*s)[i]
	return e
}

func (s *SliceStack) Top() interface{} {
	return (*s)[len(*s) - 1]
}

func (s *SliceStack) Len() int {
	return len(*s)
}

func (s *SliceStack) Empty() bool {
	return len(*s) == 0
}
