package qchan

import (
	"sync"
	"testing"
	"time"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue(3)
	if q == nil {
		t.Error("New queue should not be nil")
	}
	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}
}

func TestQueueEnqueue(t *testing.T) {
	q := NewQueue(3)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}
}

func TestQueueDequeue(t *testing.T) {
	q := NewQueue(3)
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
	q := NewQueue(3)
	_, ok := q.Dequeue()
	if ok {
		t.Error("Dequeue on empty queue should return false")
	}
}

func TestQueueIsEmpty(t *testing.T) {
	q := NewQueue(3)
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
	q := NewQueue(2)
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
	q := NewQueue(1000000)
	for i := 0; i < 1000000; i++ {
		q.Enqueue(i)
	}
	if q.Size() != 1000000 {
		t.Errorf("Expected size 1000000, got %d", q.Size())
	}
}

func TestWaitAndDequeue(t *testing.T) {
	q := NewQueue(3)

	// Test WaitAndDequeue with an empty queue
	done := make(chan struct{})
	go func() {
		time.Sleep(100 * time.Millisecond)
		q.Enqueue(1)
		close(done)
	}()

	select {
	case <-done:
		value, ok := q.WaitAndDequeue()
		if !ok || value != 1 {
			t.Errorf("Expected value 1, got %v", value)
		}
	case <-time.After(1 * time.Second):
		t.Error("Test timed out waiting for enqueue")
	}

	// Test WaitAndDequeue with multiple items
	q.Enqueue(2)
	q.Enqueue(3)

	value, ok := q.WaitAndDequeue()
	if !ok || value != 2 {
		t.Errorf("Expected value 2, got %v", value)
	}

	value, ok = q.WaitAndDequeue()
	if !ok || value != 3 {
		t.Errorf("Expected value 3, got %v", value)
	}

	// Test WaitAndDequeue with concurrent enqueue
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		q.Enqueue(4)
	}()

	select {
	case <-time.After(1 * time.Second):
		t.Error("Test timed out waiting for enqueue")
	case <-func() chan struct{} {
		done := make(chan struct{})
		go func() {
			value, ok = q.WaitAndDequeue()
			if !ok || value != 4 {
				t.Errorf("Expected value 4, got %v", value)
			}
			close(done)
		}()
		return done
	}():
	}

	wg.Wait()
}

func TestQueueConcurrentEnqueueDequeue(t *testing.T) {
	q := NewQueue(10)
	var wg sync.WaitGroup

	// Concurrently enqueue items
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			q.Enqueue(i)
		}(i)
	}

	// Concurrently dequeue items
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, ok := q.WaitAndDequeue()
			if !ok {
				t.Error("Failed to dequeue item")
			}
		}()
	}

	wg.Wait()

	if !q.IsEmpty() {
		t.Error("Queue should be empty after all operations")
	}
}
