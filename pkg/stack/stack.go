package stack

import "sync"

type Stack struct {
	items []interface{}
	lock  sync.RWMutex
}

func NewStack(input []interface{}) *Stack {
	return &Stack{items: input}
}

func (s *Stack) Push(input interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, input)
}

func (s *Stack) Pop() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = s.items[:len(s.items)-1]
}
