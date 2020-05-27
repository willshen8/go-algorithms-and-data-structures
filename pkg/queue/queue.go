package queue

import "sync"

type Queue struct {
	items []interface{}
	lock  sync.RWMutex
}

func NewQueue(input []interface{}) *Queue {
	return &Queue{items: input}
}

func (q *Queue) Enqueue(input interface{}) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = append(q.items, input)
}

func (q *Queue) Dequeue() {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = q.items[1:len(q.items)]
}
