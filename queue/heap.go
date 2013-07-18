package queue

type Comparison func(interface{}, interface{}) bool

// MakeHeap initializes a heap from a slice.
//
// Time complexity: O(n) where n is the length of a slice.
func MakeBinaryHeap(h []interface{}, less Comparison) {
	n := len(h)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(h, less, i)
	}
}

// ExtractMin retrieves minimal element of the heap.
//
// Time complexity: O(log(n)) where n is the length of the slice.
func ExtractMin(h []interface{}, less Comparison) (interface{}, []interface{}) {
	n := len(h)
	min := h[0]
	h[0] = h[n-1]
	h = h[0:(n - 1)]
	heapify(h, less, 0)
	return min, h
}

// HeapAdd appends an element to a heap.
//
// Time complexity: O(log(n)) where n = len(h)
func HeapAdd(h []interface{}, less Comparison, e interface{}) []interface{} {
	h = append(h, e)
	up(h, less, len(h)-1)
	return h
}

// PeekMin peeks minimal element of a heap.
//
// Time complexity: O(1).
func PeekMin(h []interface{}) interface{} {
	return h[0]
}

func heapify(h []interface{}, less Comparison, i int) {
	n := len(h)
	limit := n / 2
	for i < limit {
		left := 2*i + 1
		right := left + 1
		min := i

		if left < n && less(h[left], h[min]) {
			min = left
		}
		if right < n && less(h[right], h[min]) {
			min = right
		}
		if min != i {
			h[i], h[min] = h[min], h[i]
			i = min
			continue
		}
		break
	}
}

func up(h []interface{}, less Comparison, i int) {
	for {
		parent := (i - 1) / 2
		if parent == i || !less(h[i], h[parent]) {
			break
		}
		h[i], h[parent] = h[parent], h[i]
		i = parent
	}
}
