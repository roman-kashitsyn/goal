package queue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	q := NewPriorityQueue(intLess)
	const max = 10
	for i := max; i > 0; i-- {
		q.Enqueue(i)
	}
	for i := 1; i <= max; i++ {
		if k := q.Dequeue().(int); k != i {
			t.Fatalf("Invalid items order: expected %d, got %d", i, k)
		}
	}
	if !q.Empty() {
		t.Fatal("Queue expected to be empty")
	}
}
