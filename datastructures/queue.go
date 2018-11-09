package datastructures

import (
	"github.com/radlinskii/go-playground/datastructures/lists/singlylinked"
)

// Queue struct is an implementation of queue datta structure
// using singly linked list.
type Queue struct {
	list *singlylinked.List
}

// MakeQueue returns new empty queue.
func MakeQueue() *Queue {
	return &Queue{list: singlylinked.MakeList()}
}

// Enqueue ads an element at the of the queue.
func (q *Queue) Enqueue(v int) {
	q.list.Append(v)
}

// Dequeue removes an element from the front of the queue and returns it.
func (q *Queue) Dequeue() *singlylinked.Node {
	return q.list.Pop()
}

// IsEmpty checks if a queue is empty.
func (q *Queue) IsEmpty() bool {
	return q.list.IsEmpty()
}

// Front returns first element of tthe queue withour removing it.
func (q *Queue) Front() *singlylinked.Node {
	return q.list.GetHead()
}

func (q *Queue) String() string {
	if q.list.IsEmpty() {
		return "Empty queue!"
	}
	return q.list.String()
}
