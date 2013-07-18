package queue

type Queue interface {
	Peek() interface{}
	Enqueue(e interface{})
	Dequeue() interface{}
	Len() int
	Empty() bool
}

type node struct {
	item interface{}
	next *node
}

// LinkedQueue is an implementation of a Queue data structure based on
// singly-linked list. Time complexity of all the basic operations is
// O(1) in the worst case.
type LinkedQueue struct {
	head, tail *node
	len int
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{nil, nil, 0}
}

// Peek returns item which is located on the top of a queue.
//
// Time complexity: O(1)
func (q *LinkedQueue) Peek() interface{} {
	return q.head.item
}

// Enqueue appends an item to the end of a queue.
//
// Time complexity: O(1)
func (q *LinkedQueue) Enqueue(e interface{}) {
	n := &node{e, nil}
	q.len++
	if q.tail == nil {
		q.tail = n
		q.head = n
		return
	}
	q.tail.next = n
	q.tail = n
}

// Dequeue removes item on the top of a queue and returns it.
//
// Time complexity: O(1)
func (q *LinkedQueue) Dequeue() interface{} {
	h := q.head.item
	q.len--
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
		return h
	}
	q.head = q.head.next
	return h

}

// Len returns number of queued elements.
//
// Time complexity: O(1)
func (q *LinkedQueue) Len() int {
	return q.len
}

// Empty returns true if queue has no queued items.
//
// Time complexity: O(1)
func (q *LinkedQueue) Empty() bool {
	return q.tail == nil
}
