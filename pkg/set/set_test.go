package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	expected := "{}"
	actual := NewSet().String()
	assert.Equal(t, expected, actual)
}

func TestNewSetFromSlice(t *testing.T) {
	var newSetFromSliceTestCases = []struct {
		name     string
		input    []interface{}
		expected string
	}{
		{
			name:     "a slice of strings",
			input:    []interface{}{"c", "b", "a"},
			expected: "{a, b, c}",
		},
		{
			name:     "a slice of ints",
			input:    []interface{}{3, 2, 1},
			expected: "{1, 2, 3}",
		},
	}

	for _, tc := range newSetFromSliceTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewSetFromSlice(tc.input).String()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	var isEmptyTestCases = []struct {
		name     string
		input    []interface{}
		expected bool
	}{
		{
			name:     "empty set",
			input:    []interface{}{},
			expected: true,
		},
		{
			name:     "non empty set",
			input:    []interface{}{1, 2, 3},
			expected: false,
		},
	}
	for _, tc := range isEmptyTestCases {
		t.Run(tc.name, func(t *testing.T) {
			newSet := NewSetFromSlice(tc.input)
			actual := newSet.IsEmpty()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestHas(t *testing.T) {
	var hasTestCases = []struct {
		name     string
		set      []interface{}
		has      interface{}
		expected bool
	}{
		{
			name:     "don't have",
			set:      []interface{}{1, 2, 3},
			has:      1,
			expected: true,
		},
		{
			name:     "don't have",
			set:      []interface{}{1, 2, 3},
			has:      4,
			expected: false,
		},
	}
	for _, tc := range hasTestCases {
		t.Run(tc.name, func(t *testing.T) {
			newSet := NewSetFromSlice(tc.set)
			actual := newSet.Has(tc.has)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
