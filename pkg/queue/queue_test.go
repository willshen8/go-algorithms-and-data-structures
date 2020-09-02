package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	for _, tc := range newQueueTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewQueue()
			assert.Equal(t, tc.expected, actual.items)
		})
	}
}

func TestEnqueue(t *testing.T) {
	for _, tc := range enqueueTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewQueue()
			for _, action := range tc.actions {
				action(t, actual)
			}
			assert.Equal(t, tc.expected, actual.items)
		})
	}
}

func TestDeQueue(t *testing.T) {
	for _, tc := range dequeueTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewQueue()
			for _, action := range tc.actions {
				action(t, actual)
			}
			assert.Equal(t, tc.expected, actual.items)
		})
	}
}
