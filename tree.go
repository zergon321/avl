package avl

import "golang.org/x/exp/constraints"

// Tree is the implementation of
// the AVL tree data structure.
type Tree[TKey constraints.Ordered, TValue any] struct {
	root *node[TKey, TValue]
}

// Insert places the key with the value
// in the tree (overwrites the previous
// value for the key if it exists).
func (tree *Tree[TKey, TValue]) Insert(key TKey, value TValue) {
	tree.root = tree.root.insert(key, value)
}

// Search searches for the key and returns its
// value. The second returned value indicates
// if the node with the key exists in the tree.
func (tree *Tree[TKey, TValue]) Search(key TKey) (TValue, bool) {
	return tree.root.search(key)
}

// Traverse traverses the tree from the
// min key node to the max key node.
func (tree *Tree[TKey, TValue]) Traverse(visit func(currentKey TKey, currentValue TValue) error) error {
	return tree.root.traverse(visit)
}

// Remove removes the node with the specified key.
func (tree *Tree[TKey, TValue]) Remove(key TKey) {
	tree.root = tree.root.remove(key)
}

// NewTree creates a new empty tree.
func NewTree[TKey constraints.Ordered, TValue any]() *Tree[TKey, TValue] {
	return &Tree[TKey, TValue]{}
}
