package queuechan

import "testing"

package queuechan

import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	if q == nil {
		t.Error("New queue should not be nil")
	}
	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}
}

func TestQueueEnqueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}
}

func TestQueueDequeue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	item, ok := q.Dequeue()
	if !ok || item != 1 {
		t.Errorf("Expected 1, got %d", item)
	}

	item, ok = q.Dequeue()
	if !ok || item != 2 {
		t.Errorf("Expected 2, got %d", item)
	}

	item, ok = q.Dequeue()
	if !ok || item != 3 {
		t.Errorf("Expected 3, got %d", item)
	}

	if !q.IsEmpty() {
		t.Error("Queue should be empty after dequeueing all items")
	}
}

func TestQueueDequeueEmpty(t *testing.T) {
	q := NewQueue()
	_, ok := q.Dequeue()
	if ok {
		t.Error("Dequeue on empty queue should return false")
	}
}

func TestQueueIsEmpty(t *testing.T) {
	q := NewQueue()
	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}
	q.Enqueue(1)
	if q.IsEmpty() {
		t.Error("Queue should not be empty after enqueue")
	}
	q.Dequeue()
	if !q.IsEmpty() {
		t.Error("Queue should be empty after dequeueing all items")
	}
}

func TestQueueSize(t *testing.T) {
	q := NewQueue()
	if q.Size() != 0 {
		t.Errorf("Expected size 0, got %d", q.Size())
	}
	q.Enqueue(1)
	if q.Size() != 1 {
		t.Errorf("Expected size 1, got %d", q.Size())
	}
	q.Enqueue(2)
	if q.Size() != 2 {
		t.Errorf("Expected size 2, got %d", q.Size())
	}
	q.Dequeue()
	if q.Size() != 1 {
		t.Errorf("Expected size 1, got %d", q.Size())
	}
	q.Dequeue()
	if q.Size() != 0 {
		t.Errorf("Expected size 0, got %d", q.Size())
	}
}
