package stack

import "testing"

// Action calls a function of the stack
type Action func(*testing.T, *Stack)

var newStackTestCases = []struct {
	name     string
	input    []interface{}
	expected []interface{}
}{
	{
		name:     "stack with no data",
		input:    []interface{}{},
		expected: []interface{}{},
	},
	{
		name:     "stack with one item",
		input:    []interface{}{1},
		expected: []interface{}{1},
	},
	{
		name:     "stack with multiple items",
		input:    []interface{}{1, 2, 3, 4, 5},
		expected: []interface{}{1, 2, 3, 4, 5},
	},
	{
		name:     "stack with mixed items",
		input:    []interface{}{1, "two", 3, "four", 5},
		expected: []interface{}{1, "two", 3, "four", 5},
	},
}

func push(arg interface{}) Action {
	return func(t *testing.T, q *Stack) {
		q.Push(arg)
	}
}

var pushTestCases = []struct {
	name     string
	actions  []Action
	expected []interface{}
}{
	{
		name: "adding a single item into the stack",
		actions: []Action{
			push(1),
		},
		expected: []interface{}{1},
	},
	{
		name: "adding multiple items into the stack",
		actions: []Action{
			push(1),
			push(2),
			push(3),
			push(4),
			push(5),
		},
		expected: []interface{}{1, 2, 3, 4, 5},
	},
}

func pop() Action {
	return func(t *testing.T, q *Stack) {
		q.Pop()
	}
}

var popTestCases = []struct {
	name     string
	actions  []Action
	expected []interface{}
}{
	{
		name: "remove a single item from stack",
		actions: []Action{
			push(1),
			pop(),
		},
		expected: []interface{}{},
	},
	{
		name: "adding multiple items into the stack",
		actions: []Action{
			push(1),
			push(2),
			pop(),
		},
		expected: []interface{}{1},
	},
}
