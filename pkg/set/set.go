package set

import (
	"fmt"
	"reflect"
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

// Subset returns true if s1 is a subset of s2, e.g. all elements of s1 exists in s2
func Subset(s1, s2 Set) bool {
	for k := range s1 {
		if !s2.Has(k) {
			return false
		}
	}
	return true
}

// Disjoint return true if there are no common elements between two sets
func Disjoint(s1, s2 Set) bool {
	for k := range s1 {
		if s2.Has(k) {
			return false
		}
	}
	return true
}

// Equal return true if s1 and s2 are exactly the same
func Equal(s1, s2 Set) bool {
	return reflect.DeepEqual(s1, s2)
}

// Intersection returns a set with all the common elements between the two sets
func Intersection(s1, s2 Set) Set {
	resultSet := NewSet()
	for key := range s1 {
		if s2.Has(key) {
			resultSet[key] = true
		}
	}
	return resultSet
}

// Difference returns the elements that only exists in the first set
func Difference(s1, s2 Set) Set {
	resultSet := NewSet()
	for key := range s1 {
		if !s2.Has(key) {
			resultSet[key] = true
		}
	}
	return resultSet
}

// Union returns a set that has all the elements in set 1 and set 2
func Union(s1, s2 Set) Set {
	unionSet := NewSet()
	for key := range s1 { // copy s1 into resultSet
		unionSet[key] = true
	}
	for key := range s2 {
		if !unionSet.Has(key) {
			unionSet[key] = true
		}
	}
	return unionSet
}
