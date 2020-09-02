package queue

import "testing"

// Action calls a function of the queue
type Action func(*testing.T, *Queue)

var newQueueTestCases = []struct {
	name     string
	expected []interface{}
}{
	{
		name:     "queue with no data",
		expected: []interface{}(nil),
	},
}

func enqueue(arg interface{}) Action {
	return func(t *testing.T, q *Queue) {
		q.Enqueue(arg)
	}
}

var enqueueTestCases = []struct {
	name     string
	actions  []Action
	expected []interface{}
}{
	{
		name: "adding a single item into the queue",
		actions: []Action{
			enqueue(1),
		},
		expected: []interface{}{1},
	},
	{
		name: "adding multiple items into the queue",
		actions: []Action{
			enqueue(1),
			enqueue(2),
			enqueue(3),
			enqueue(4),
			enqueue(5),
		},
		expected: []interface{}{1, 2, 3, 4, 5},
	},
}

func dequeue() Action {
	return func(t *testing.T, q *Queue) {
		q.Dequeue()
	}
}

var dequeueTestCases = []struct {
	name     string
	actions  []Action
	expected []interface{}
}{
	{
		name: "remove a single item from queue",
		actions: []Action{
			enqueue(1),
			dequeue(),
		},
		expected: []interface{}{},
	},
	{
		name: "adding multiple items into the queue",
		actions: []Action{
			enqueue(1),
			enqueue(2),
			dequeue(),
		},
		expected: []interface{}{2},
	},
}