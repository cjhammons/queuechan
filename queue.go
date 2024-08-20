package queuechan

import "sync"

// Node represents a node in the queue.
//
// Value is the value of the node.
// Next is the next node in the queue.
type Node struct {
	Value any
	Next  *Node
}

// Queue represents a queue.
//
// head is the head of the queue.
// tail is the tail of the queue.
// lock is a mutex to lock the queue for exclusive access.
// size is the size of the queue.
// ch is a channel to signal when an item is available.
type Queue struct {
	head *Node
	tail *Node
	lock sync.Mutex
	size int
	ch   chan struct{}
}

// NewQueue creates and returns a new Queue.
func NewQueue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
		lock: sync.Mutex{},
		ch:   make(chan struct{}),
	}
}

// Enqueue enqueues an item to the queue.
func (q *Queue) Enqueue(value any) {
	q.lock.Lock()
	defer q.lock.Unlock()

	q.size++

	if q.head == nil {
		q.head = &Node{Value: value}
		q.tail = q.head
	} else {
		q.tail.Next = &Node{Value: value}
		q.tail = q.tail.Next
	}
}

// WaitAndDequeue waits for an item to be available and then dequeues it.
// WARNING: Untested
func (q *Queue) WaitAndDequeue() (any, bool) {
	<-q.ch // wait for the signal that an item is available
	return q.Dequeue()
}

// Dequeue dequeues an item from the queue.
//
// Returns the value of the item and a boolean indicating if the item was successfully dequeued.
func (q *Queue) Dequeue() (any, bool) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.head == nil {
		return nil, false
	}

	q.size--
	value := q.head.Value
	q.head = q.head.Next

	return value, true
}

// IsEmpty returns true if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// Peak returns the value of the head of the queue and a boolean indicating if the head was successfully returned.
func (q *Queue) Peak() (any, bool) {
	if q.head == nil {
		return nil, false
	}

	return q.head.Value, true
}

// Size returns the size of the queue.
func (q *Queue) Size() int {
	return q.size
}
