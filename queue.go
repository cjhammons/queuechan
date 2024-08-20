package qchan

import "sync"

// Queue represents a queue using a channel.
// It is safe for concurrent use.
//
// ch is the underlying channel that stores the queue elements.
// lock is a mutex used to synchronize access to the queue.
// size is the current number of elements in the queue.
type Queue struct {
	ch   chan any
	lock sync.Mutex
	size int
}

// NewQueue creates and returns a new Queue with a specified buffer size.
//
// bufferSize is the size of the channel buffer.
func NewQueue(bufferSize int) *Queue {
	return &Queue{
		ch: make(chan any, bufferSize),
	}
}

// Enqueue adds a new value to the end of the queue.
//
// value is the value to be added to the queue.
func (q *Queue) Enqueue(value any) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.ch <- value
	q.size++
}

// WaitAndDequeue waits for an item to be available and then dequeues it.
// Returns the dequeued value and a boolean indicating if the operation was successful.
func (q *Queue) WaitAndDequeue() (any, bool) {
	value, ok := <-q.ch
	if ok {
		q.lock.Lock()
		q.size--
		q.lock.Unlock()
	}
	return value, ok
}

// Dequeue removes and returns the value from the front of the queue.
// Returns the dequeued value and a boolean indicating if the operation was successful.
func (q *Queue) Dequeue() (any, bool) {
	select {
	case value := <-q.ch:
		q.lock.Lock()
		q.size--
		q.lock.Unlock()
		return value, true
	default:
		return nil, false
	}
}

// IsEmpty checks if the queue is empty.
// Returns true if the queue is empty, otherwise false.
func (q *Queue) IsEmpty() bool {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.size == 0
}

// Size returns the number of items in the queue.
func (q *Queue) Size() int {
	q.lock.Lock()
	defer q.lock.Unlock()
	return q.size
}
