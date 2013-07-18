package stack

import (
	"testing"
	"testing/quick"
)

func checkStackOrder(ints []int) bool {
	n := len(ints)
	s := NewSliceStackOfCapacity(n)
	for _, i := range ints {
		s.Push(i)
	}
	for i := n - 1; i >= 0; i-- {
		if e := s.Pop(); e != ints[i] {
			return false
		}
	}
	return s.Empty()

}

func TestStackProperty(t *testing.T) {
	if err := quick.Check(checkStackOrder, nil); err != nil {
		t.Error(err)
	}
}
