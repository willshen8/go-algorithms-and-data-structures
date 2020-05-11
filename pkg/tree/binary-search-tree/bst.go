package bst

import (
	"fmt"
	"sync"
)

// Node in a binary search tree contains a key and value and reference to left and right child
type Node struct {
	key       int
	value     int
	leftNode  *Node
	rightNode *Node
}

type BST struct {
	root *Node
	lock sync.RWMutex
}

func (n Node) String() string {
	return fmt.Sprintf("Key: %d, Value: %d, LeftNode: %s, RightNode: %s", n.key, n.value, n.leftNode, n.rightNode)
}

// Insert will add new value into the tree
func (tree *BST) Insert(key int, value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	newNode := &Node{key, value, nil, nil}

	if tree.root == nil {
		tree.root = newNode
	} else {
		InsertNode(tree.root, newNode)
	}
}

// InsertNode will insert a node starting from the root
func InsertNode(root *Node, node *Node) {
	if node.key < root.key {
		if root.leftNode == nil {
			root.leftNode = node
		} else {
			InsertNode(root.leftNode, node)
		}
	} else if node.key > root.key {
		if root.rightNode == nil {
			root.rightNode = node
		} else {
			InsertNode(root.rightNode, node)
		}
	} else { //case when key equals to root key
		root.value = node.value
	}
}

// Search will look for a particular key in the tree and return a boolean
func (tree *BST) SearchNode(key int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	return searchNode(tree.root, key)
}

func searchNode(root *Node, key int) bool {
	if root == nil {
		return false
	} else if key < root.key {
		return searchNode(root.leftNode, key)
	} else if key > root.key {
		return searchNode(root.rightNode, key)
	}
	return true
}

// RemoveNode removes a node with a key from the BST
func (tree *BST) DeleteNode(key int) *Node {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	return deleteNode(tree.root, key)
}

func deleteNode(root *Node, key int) *Node {
	if root == nil {
		return nil
	} else if root.key < key {
		return deleteNode(root.leftNode, key)
	} else if root.key > key {
		return deleteNode(root.rightNode, key)
	}
	// when key == root.
	return nil
}
