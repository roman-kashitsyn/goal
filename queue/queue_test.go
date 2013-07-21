package queue

import (
	"testing"
)

func TestQueueOrder(t *testing.T) {
	qc(t, func(ints []int) bool {
		q := NewLinkedQueue()
		for _, i := range ints {
			q.Enqueue(i)
		}
		for _, i := range ints {
			if q.Peek() != i || q.Dequeue() != i {
				return false
			}
		}
		return q.Empty()
	})
}

func TestEmptyQueue(t *testing.T) {
	q := NewLinkedQueue()
	if !q.Empty() {
		t.Fatal("New linked queue expected to be Empty()")
	}
}
