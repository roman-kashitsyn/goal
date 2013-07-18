package queue

const initialCapacity = 10

type PriorityQueue struct {
	heap []interface{}
	less Comparison
}

func NewPriorityQueue(less Comparison) *PriorityQueue {
	return &PriorityQueue{heap: make([]interface{}, 0, initialCapacity), less: less}
}

func (q *PriorityQueue) Enqueue(item interface{}) {
	q.heap = HeapAdd(q.heap, q.less, item)
}

func (q *PriorityQueue) Dequeue() interface{} {
	e, nq := ExtractMin(q.heap, q.less)
	q.heap = nq
	return e
}

func (q *PriorityQueue) Peek() interface{} {
	return PeekMin(q.heap)
}

func (q *PriorityQueue) Empty() bool {
	return len(q.heap) == 0
}

func (q *PriorityQueue) Len() int {
	return len(q.heap)
}
