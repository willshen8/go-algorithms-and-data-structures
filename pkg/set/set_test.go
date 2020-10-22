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

func TestSubset(t *testing.T) {
	var subsetTestCases = []struct {
		name     string
		set1     []interface{}
		set2     []interface{}
		expected bool
	}{
		{
			name:     "empty set",
			set1:     []interface{}{},
			set2:     []interface{}{},
			expected: true,
		},
		{
			name:     "s1 is subset of s2",
			set1:     []interface{}{1},
			set2:     []interface{}{1, 2, 3},
			expected: true,
		},
		{
			name:     "s1 is NOT a subset of s2",
			set1:     []interface{}{1},
			set2:     []interface{}{2, 3},
			expected: false,
		},
	}
	for _, tc := range subsetTestCases {
		t.Run(tc.name, func(t *testing.T) {
			set1 := NewSetFromSlice(tc.set1)
			set2 := NewSetFromSlice(tc.set2)
			actual := Subset(set1, set2)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestDisjoint(t *testing.T) {
	var disjointTestCases = []struct {
		name     string
		set1     []interface{}
		set2     []interface{}
		expected bool
	}{
		{
			name:     "empty set",
			set1:     []interface{}{},
			set2:     []interface{}{},
			expected: true,
		},
		{
			name:     "No common elements between s1 and s2",
			set1:     []interface{}{1},
			set2:     []interface{}{2, 3},
			expected: true,
		},
		{
			name:     "There is one common element between s1 and s2",
			set1:     []interface{}{"a"},
			set2:     []interface{}{"a", "b", "c"},
			expected: false,
		},
	}
	for _, tc := range disjointTestCases {
		t.Run(tc.name, func(t *testing.T) {
			set1 := NewSetFromSlice(tc.set1)
			set2 := NewSetFromSlice(tc.set2)
			actual := Disjoint(set1, set2)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestEqual(t *testing.T) {
	var equalTestCases = []struct {
		name     string
		set1     []interface{}
		set2     []interface{}
		expected bool
	}{
		{
			name:     "empty set are equal",
			set1:     []interface{}{},
			set2:     []interface{}{},
			expected: true,
		},
		{
			name:     "Equal",
			set1:     []interface{}{1, 2, 3},
			set2:     []interface{}{3, 2, 1},
			expected: true,
		},
		{
			name:     "Not Equal",
			set1:     []interface{}{1, 2, 3},
			set2:     []interface{}{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "Two strings set are not equal",
			set1:     []interface{}{1, 2, 3, 4},
			set2:     []interface{}{1, 2, 3},
			expected: false,
		},
	}
	for _, tc := range equalTestCases {
		t.Run(tc.name, func(t *testing.T) {
			set1 := NewSetFromSlice(tc.set1)
			set2 := NewSetFromSlice(tc.set2)
			actual := Equal(set1, set2)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestIntersection(t *testing.T) {
	var intersectionTestCases = []struct {
		name     string
		set1     []interface{}
		set2     []interface{}
		expected string
	}{
		{
			name:     "empty set have no intersections",
			set1:     []interface{}{},
			set2:     []interface{}{},
			expected: "{}",
		},
		{
			name:     "One common element",
			set1:     []interface{}{1, 2, 3},
			set2:     []interface{}{3, 4, 5},
			expected: "{3}",
		},
		{
			name:     "No intersection",
			set1:     []interface{}{1, 2, 3},
			set2:     []interface{}{4, 5, 6},
			expected: "{}",
		},
	}
	for _, tc := range intersectionTestCases {
		t.Run(tc.name, func(t *testing.T) {
			set1 := NewSetFromSlice(tc.set1)
			set2 := NewSetFromSlice(tc.set2)
			actual := Intersection(set1, set2).String()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestDifference(t *testing.T) {
	var differenceTestCases = []struct {
		name     string
		set1     []interface{}
		set2     []interface{}
		expected string
	}{
		{
			name:     "empty set have no difference",
			set1:     []interface{}{},
			set2:     []interface{}{},
			expected: "{}",
		},
		{
			name:     "two different elements",
			set1:     []interface{}{1, 2, 3},
			set2:     []interface{}{3, 4, 5},
			expected: "{1, 2}",
		},
		{
			name:     "No intersection",
			set1:     []interface{}{1, 2, 3},
			set2:     []interface{}{1, 2},
			expected: "{3}",
		},
	}
	for _, tc := range differenceTestCases {
		t.Run(tc.name, func(t *testing.T) {
			set1 := NewSetFromSlice(tc.set1)
			set2 := NewSetFromSlice(tc.set2)
			actual := Difference(set1, set2).String()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestUnion(t *testing.T) {
	var unionTestCases = []struct {
		name     string
		set1     []interface{}
		set2     []interface{}
		expected string
	}{
		{
			name:     "empty set have no difference",
			set1:     []interface{}{},
			set2:     []interface{}{},
			expected: "{}",
		},
		{
			name:     "Only one duplicate element",
			set1:     []interface{}{1, 2, 3},
			set2:     []interface{}{3, 4, 5},
			expected: "{1, 2, 3, 4, 5}",
		},
		{
			name:     "First set covers everything",
			set1:     []interface{}{1, 2, 3},
			set2:     []interface{}{1, 2},
			expected: "{1, 2, 3}",
		},
	}
	for _, tc := range unionTestCases {
		t.Run(tc.name, func(t *testing.T) {
			set1 := NewSetFromSlice(tc.set1)
			set2 := NewSetFromSlice(tc.set2)
			actual := Union(set1, set2).String()
			assert.Equal(t, tc.expected, actual)
		})
	}
}
