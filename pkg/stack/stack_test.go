package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStack(t *testing.T) {
	for _, tc := range newStackTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewStack(tc.input)
			assert.Equal(t, tc.expected, actual.items)
		})
	}
}

func TestPush(t *testing.T) {
	for _, tc := range pushTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewStack([]interface{}{})
			for _, action := range tc.actions {
				action(t, actual)
			}
			assert.Equal(t, tc.expected, actual.items)
		})
	}
}

func TestPop(t *testing.T) {
	for _, tc := range popTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewStack([]interface{}{})
			for _, action := range tc.actions {
				action(t, actual)
			}
			assert.Equal(t, tc.expected, actual.items)
		})
	}
}
