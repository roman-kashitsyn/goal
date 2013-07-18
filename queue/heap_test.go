package queue

import (
	"math"
	"testing"
	"testing/quick"
)

func intLess(x, y interface{}) bool {
	return x.(int) < y.(int)
}

func toGenericSlice(ints []int) []interface{} {
	h := make([]interface{}, len(ints))
	for i, x := range ints {
		h[i] = x
	}
	return h
}

func intOrdered(h []interface{}) bool {
	n := len(h)
	prev := math.MinInt32
	var m interface{}
	for i := 0; i < n; i++ {
		m, h = ExtractMin(h, intLess)
		if !intLess(prev, m) {
			return false
		}
		prev = m.(int)
	}
	return true
}

func isHeap(h []interface{}) bool {
	n := len(h)
	for i := 0; i < n/2; i++ {
		l := 2*i + 1
		r := l + 1
		if l < n && intLess(h[l], h[i]) {
			return false
		}
		if r < n && intLess(h[r], h[i]) {
			return false
		}
	}
	return true
}

func checkHeapProperty(ints []int) bool {
	h := toGenericSlice(ints)
	MakeBinaryHeap(h, intLess)
	return isHeap(h)
}

func TestHeapOrder(t *testing.T) {
	if err := quick.Check(checkHeapProperty, nil); err != nil {
		t.Error(err)
	}
}

func checkMinProperty(ints []int) bool {
	h := toGenericSlice(ints)
	MakeBinaryHeap(h, intLess)
	return intOrdered(h)
}

func TestMinExtraction(t *testing.T) {
	if err := quick.Check(checkMinProperty, nil); err != nil {
		t.Error(err)
	}
}

func checkIncrementalConstruction(ints []int) bool {
	h := make([]interface{}, 0, len(ints))
	for _, i := range ints {
		h = HeapAdd(h, intLess, i)
		if !isHeap(h) {
			return false
		}
	}
	return true
}

func TestIncrementalConstruction(t *testing.T) {
	if err := quick.Check(checkIncrementalConstruction, nil); err != nil {
		t.Error(err)
	}
}
