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
	expectedRightNode := &Node{key: 2, value: 2, leftNode: nil, rightNode: nil}
	if !reflect.DeepEqual(tree.root.rightNode, expectedRightNode) {
		t.Fatalf("Test case failed: Insert a node into a tree with a single root")
	}
}

func TestInsertRightLeft(t *testing.T) {
	tree := &BST{}
	tree.Insert(1, 1)
	tree.Insert(0, 0)
	tree.Insert(2, 2)
	expectedLeftNode := &Node{key: 0, value: 0, leftNode: nil, rightNode: nil}
	if !reflect.DeepEqual(tree.root.leftNode, expectedLeftNode) {
		t.Fatalf("Test case failed: Insert 3 nodes into a tree")
	}
}

func TestInsertNodeEqualKey(t *testing.T) {
	tree := &BST{}
	tree.Insert(1, 1)
	tree.Insert(1, 0)
	expectedRoot := &Node{key: 1, value: 0, leftNode: nil, rightNode: nil}
	if !reflect.DeepEqual(tree.root, expectedRoot) {
		t.Fatalf("Test case failed: Insert a nodes with same key but different value")
	}
}

func TestSearchRootSuccess(t *testing.T) {
	tree := &BST{}
	tree.Insert(0, 0)
	expected := true
	actual := tree.SearchNode(0)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Test case failed: Expected result: %t, Actual result: %t", expected, actual)
	}
}

func TestSearchNodeFail(t *testing.T) {
	tree := &BST{}
	tree.Insert(0, 0)
	expected := false
	actual := tree.SearchNode(1)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Test case failed: Expected result: %t, Actual result: %t", expected, actual)
	}
}

// func TestDeleteRoot(t *testing.T) {
// 	tree := &BST{}
// 	tree.Insert(0, 0)
// 	expected := &Node{}
// 	actual := tree.DeleteNode(0)
// 	if !reflect.DeepEqual(expected, actual) {
// 		t.Fatalf("Test case failed: Expected result: %s, Actual result: %s", expected, actual)
// 	}
// }
