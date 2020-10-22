package set

import (
	"fmt"
	"sort"
)

type Set map[interface{}]bool

func NewSet() Set {
	return make(map[interface{}]bool)
}

// String prints out the set in sorted order
func (s Set) String() string {
	var sortedSet []string
	for k := range s {
		sortedSet = append(sortedSet, fmt.Sprintf("%v", k))
	}

	sort.Strings(sortedSet)
	result := "{"
	for i, v := range sortedSet {
		result += v
		if i != len(s)-1 {
			result += ", "
		}
	}
	result += "}"
	return result
}

// NewSetFromSlice creates a new set from a slice of input
func NewSetFromSlice(input []interface{}) Set {
	var newSet = make(map[interface{}]bool)
	for _, v := range input {
		if _, ok := newSet[v]; !ok {
			newSet[v] = true
		}
	}
	return newSet
}

// IsEmpty returns true if the set is empty
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Has return a true if a value exists in a set
func (s Set) Has(value interface{}) bool {
	for key := range s {
		if key == value {
			return true
		}
	}
	return false
}
