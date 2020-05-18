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

func TestInsertTwoLeafNodes(t *testing.T) {
	tree := &BST{}
	tree.Insert(1, 1)
	tree.Insert(0, 0)
	tree.Insert(2, 2)
	expected := &Node{key: 1, value: 1, leftNode: &Node{key: 0, value: 0}, rightNode: &Node{key: 2, value: 2}}
	if !reflect.DeepEqual(tree.root, expected) {
		t.Fatalf("Test case failed: Insert two leaf nodes")
	}
}

func TestInsertRigthUnbalancedTree(t *testing.T) {
	tree := &BST{}
	tree.Insert(0, 0)
	tree.Insert(1, 1)
	tree.Insert(2, 2)
	expected := &Node{key: 0, value: 0, leftNode: nil, rightNode: &Node{key: 1, value: 1, leftNode: nil, rightNode: &Node{key: 2, value: 2}}}
	if !reflect.DeepEqual(tree.root, expected) {
		t.Fatalf("Test case failed: Insert two right leaf nodes")
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

func TestSearchSuccess(t *testing.T) {
	tree := &BST{}
	tree.Insert(0, "Hello")
	tree.Insert(1, "World!")
	expected := "World!"
	actual := tree.SearchNode(1)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Test case failed: Expected result: %s, Actual result: %s", expected, actual)
	}
}

func TestSearchNodeFail(t *testing.T) {
	tree := &BST{}
	actual := tree.SearchNode(1)
	if actual != nil {
		t.Fatalf("Test case failed: Expected result: nil, Actual result: %t", actual)
	}
}

func TestDeleteEmptyTree(t *testing.T) {
	tree := &BST{}
	deleteResult := tree.DeleteNode(0)
	if deleteResult != nil {
		t.Fatalf("Test case failed: delete an empty tree should output nil, got %v", deleteResult)
	}
}

func TestDeleteWithNoNodes(t *testing.T) {
	tree := &BST{}
	tree.Insert(0, "Hello, World!")
	actual := tree.DeleteNode(0)
	if actual != nil {
		t.Fatalf("Test case failed: tree should be empty with no nodes, got %v", actual)
	}
}

func TestDeleteNodeWithNoLeftChild(t *testing.T) {
	tree := &BST{}
	tree.Insert(0, "A")
	tree.Insert(1, "B")
	tree.DeleteNode(1)
	searchResult := tree.SearchNode(1)
	if searchResult != nil {
		t.Fatalf("Test case failed: root node with key 0 should have been deleted, got %s", searchResult)
	}
	expectedTree := &Node{key: 0, value: "A", leftNode: nil, rightNode: nil}
	if !reflect.DeepEqual(tree.root, expectedTree) {
		t.Fatalf("Test case failed: Delete a right leaf node only, expected: root(0) -> right(1)")
	}
}

func TestDeleteNodeWithNoRightChild(t *testing.T) {
	tree := &BST{}
	tree.Insert(1, "A")
	tree.Insert(0, "B")
	tree.DeleteNode(0)
	searchResult := tree.SearchNode(0)
	if searchResult != nil {
		t.Fatalf("Test case failed: root node with key 0 should have been deleted, got %s", searchResult)
	}
}

func TestDeleteRootWithNoLeftChild(t *testing.T) {
	tree := &BST{}
	tree.Insert(1, "A")
	tree.Insert(2, "B")
	actual := tree.DeleteNode(1)
	expectedTree := &Node{key: 2, value: "B", leftNode: nil, rightNode: nil}
	if !reflect.DeepEqual(actual, expectedTree) {
		t.Fatalf("Test case failed: Delete a node with one child, expected: root(2)")
	}
}

func TestDeleteLeftLeafNode(t *testing.T) {
	tree := &BST{}
	tree.Insert(1, "B")
	tree.Insert(0, "A")
	tree.Insert(2, "C")
	tree.DeleteNode(0)
	searchResult := tree.SearchNode(0)
	if searchResult != nil {
		t.Fatalf("Test case failed: node with key 0 should have been deleted, got %s", searchResult)
	}
	expectedTree := &Node{key: 1, value: "B", leftNode: nil, rightNode: &Node{key: 2, value: "C"}}
	if !reflect.DeepEqual(tree.root, expectedTree) {
		t.Fatalf("Test case failed: Delete a node with one child, expected: root(2) -> left(0)")
	}
}

func TestDeleteRightLeafNode(t *testing.T) {
	tree := &BST{}
	tree.Insert(0, "A")
	tree.Insert(1, "B")
	tree.Insert(2, "C")
	tree.DeleteNode(2)
	searchResult := tree.SearchNode(2)
	if searchResult != nil {
		t.Fatalf("Test case failed: node with key 2 should have been deleted, got %s", searchResult)
	}
	expectedTree := &Node{key: 0, value: "A", leftNode: nil, rightNode: &Node{key: 1, value: "B"}}
	if !reflect.DeepEqual(tree.root, expectedTree) {
		t.Fatalf("Test case failed: Delete a right leaf node, expected: root(0) -> left(1)")
	}
}

func TestDeleteNodeWithOneChild(t *testing.T) {
	tree := &BST{}
	tree.Insert(2, 2)
	tree.Insert(1, 1)
	tree.Insert(0, 0)
	tree.DeleteNode(1)
	expectedTree := &Node{key: 2, value: 2, leftNode: &Node{key: 0, value: 0}, rightNode: nil}
	if !reflect.DeepEqual(tree.root, expectedTree) {
		t.Fatalf("Test case failed: Delete a node with one child, expected: root(2) -> left(0)")
	}
}

func TestDeleteNodeWithTwoChildren(t *testing.T) {
	tree := &BST{}
	tree.Insert(1, "B")
	tree.Insert(0, "A")
	tree.Insert(2, "C")
	tree.DeleteNode(1)
	searchResult := tree.SearchNode(1)
	if searchResult != nil {
		t.Fatalf("Test case failed: root node with key 1 should have been deleted, got %s", searchResult)
	}
	expectedTree := &Node{key: 2, value: "C", leftNode: &Node{key: 0, value: "A", rightNode: nil}}
	if !reflect.DeepEqual(tree.root, expectedTree) {
		t.Fatalf("Test case failed: Delete a node with 2 children, expected: root(0) -> right(2)")
	}
}
