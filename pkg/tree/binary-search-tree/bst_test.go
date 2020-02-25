package bst

import (
	"reflect"
	"testing"
)

func TestInsertEmptyTree(t *testing.T) {
	tree := &BST{}
	tree.Insert(1, 1)
	expected := &BST{root: &Node{key: 1, value: 1, leftNode: nil, rightNode: nil}}
	if !reflect.DeepEqual(tree, expected) {
		t.Fatalf("Test case failed : Insert a node into empty tree")
	}
}

func TestInsertRight(t *testing.T) {
	tree := &BST{}
	tree.Insert(1, 1)
	tree.Insert(2, 2)
	expected := &BST{root: &Node{key: 1, value: 1, leftNode: nil, rightNode: &Node{key: 2, value: 2, leftNode: nil, rightNode: nil}}}
	if !reflect.DeepEqual(tree, expected) {
		t.Fatalf("Test case failed: Insert a node into a tree with a single root")
	}
}

func TestInsertRightLeft(t *testing.T) {
	tree := &BST{}
	tree.Insert(1, 1)
	tree.Insert(0, 0)
	tree.Insert(2, 2)
	expected := &BST{root: &Node{key: 1, value: 1, leftNode: &Node{key: 0, value: 0, leftNode: nil, rightNode: nil}, rightNode: &Node{key: 2, value: 2, leftNode: nil}}}
	if !reflect.DeepEqual(tree, expected) {
		t.Fatalf("Test case failed: Insert 3 nodes into a tree")
	}
}
