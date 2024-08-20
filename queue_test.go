package qchan

import (
	"sync"
	"testing"
	"time"
)

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

func TestLargeQueue(t *testing.T) {
	q := NewQueue()
	for i := 0; i < 1000000; i++ {
		q.Enqueue(i)
	}
	if q.Size() != 1000000 {
		t.Errorf("Expected size 1000000, got %d", q.Size())
	}
}

func TestQueueConcurrentEnqueue(t *testing.T) {
	q := NewQueue()
	var wg sync.WaitGroup
	numGoroutines := 100
	numItems := 1000

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func(start int) {
			defer wg.Done()
			for j := 0; j < numItems; j++ {
				q.Enqueue(start*numItems + j)
			}
		}(i)
	}
	wg.Wait()

	expectedSize := numGoroutines * numItems
	if q.Size() != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, q.Size())
	}
}

func TestQueueConcurrentDequeue(t *testing.T) {
	q := NewQueue()
	numGoroutines := 100
	numItems := 1000

	// Enqueue items
	for i := 0; i < numGoroutines*numItems; i++ {
		q.Enqueue(i)
	}

	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numItems; j++ {
				q.Dequeue()
			}
		}()
	}
	wg.Wait()

	if !q.IsEmpty() {
		t.Error("Queue should be empty after concurrent dequeues")
	}
}
