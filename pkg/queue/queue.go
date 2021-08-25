package queue

import "sync"

// Queue is a structure that contains a slice of any type and a mutex
type Queue struct {
	items []interface{}
	lock  sync.RWMutex
}

func NewQueue() *Queue {
	return &Queue{}
}

// Enqueue puts a new item at the back of the queue
func (q *Queue) Enqueue(input interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = append(q.items, input)
}

// Dequeue take one item off the queue
func (q *Queue) Dequeue() interface{} {
	q.lock.Lock()
	defer q.lock.Unlock()
	dequeuedItem := q.items[0]
	q.items = q.items[1:len(q.items)]
	return dequeuedItem
}

// Length gives the length of the queue
func (q *Queue) Length() int {
	return len(q.items)
}
