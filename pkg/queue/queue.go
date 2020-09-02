package queue

import "sync"

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
func (q *Queue) Dequeue() {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = q.items[1:len(q.items)]
}

// Length gives the length of the queue
func (q *Queue) Length() int {
	return len(q.items)
}
