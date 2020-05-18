package bst

import (
	"sync"
)

// Node in a binary search tree contains a key and value and reference to left and right child
type Node struct {
	key       int
	value     interface{}
	leftNode  *Node
	rightNode *Node
}

// BST contains a link to its root and a mutex to synchronise write and read operations
type BST struct {
	root *Node
	lock sync.RWMutex
}

// Insert will add new value into the tree
func (tree *BST) Insert(key int, value interface{}) {
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

// SearchNode will look for a particular key in the tree and return a boolean
func (tree *BST) SearchNode(key int) interface{} {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	return searchNode(tree.root, key)
}

func searchNode(root *Node, key int) interface{} {
	if root == nil {
		return nil
	} else if key < root.key {
		return searchNode(root.leftNode, key)
	} else if key > root.key {
		return searchNode(root.rightNode, key)
	}
	return root.value
}

// DeleteNode removes a node from the BST and returns the updated tree
func (tree *BST) DeleteNode(key int) *Node {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	return deleteNode(tree.root, key)
}

func deleteNode(root *Node, key int) *Node {
	if root == nil {
		return nil
	}
	if key < root.key {
		root.leftNode = deleteNode(root.leftNode, key)
		return root
	}
	if key > root.key {
		root.rightNode = deleteNode(root.rightNode, key)
		return root
	}
	// key == root.key
	if root.leftNode == nil && root.rightNode == nil {
		root = nil
		return root
	}
	if root.leftNode == nil { // make right tree the new tree
		root = root.rightNode
		return root
	}
	if root.rightNode == nil { // make left tree the new tree
		root = root.leftNode
		return root
	}
	// Find the left most right node for the replacement
	leftMostRightNode := root.rightNode
	for {
		//find smallest value on the right side
		if leftMostRightNode != nil && leftMostRightNode.leftNode != nil {
			leftMostRightNode = leftMostRightNode.leftNode
		} else {
			break
		}
	}
	root.key, root.value = leftMostRightNode.key, leftMostRightNode.value
	root.rightNode = deleteNode(root.rightNode, root.key)
	return root
}
