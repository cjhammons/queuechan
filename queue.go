package queuechan

import "sync"

type Node struct {
	Value any
	Next  *Node
}

type Queue struct {
	head *Node
	tail *Node
	lock sync.Mutex
	size int
	ch   chan struct{}
}

func NewQueue() *Queue {
	return &Queue{
		head: nil,
		tail: nil,
		lock: sync.Mutex{},
		ch:   make(chan struct{}),
	}
}

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

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Peak() (any, bool) {
	if q.head == nil {
		return nil, false
	}

	return q.head.Value, true
}

func (q *Queue) Size() int {
	return q.size
}
