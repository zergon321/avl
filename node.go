package avl

import (
	"golang.org/x/exp/constraints"
)

// node is the representation
// of the AVL tree node.
type node[TKey constraints.Ordered, TValue any] struct {
	height int
	left   *node[TKey, TValue]
	right  *node[TKey, TValue]
	key    TKey
	value  TValue
}

// getHeight returns the level of
// the node in the AVL tree.
func (nd *node[TKey, TValue]) getHeight() int {
	if nd == nil {
		return 0
	}

	return nd.height
}

// balanceFactor returns the difference
// between subtrees heights.
func (nd *node[TKey, TValue]) balanceFactor() int {
	return nd.right.getHeight() - nd.left.getHeight()
}

// fixHeight assigns a correct
// height to the node.
func (nd *node[TKey, TValue]) fixHeight() {
	leftHeight := nd.left.getHeight()
	rightHeight := nd.right.getHeight()

	if leftHeight > rightHeight {
		nd.height = leftHeight + 1
	} else {
		nd.height = rightHeight + 1
	}
}

// rotateRight performs the right
// rotation in order to balance the tree.
func (nd *node[TKey, TValue]) rotateRight() *node[TKey, TValue] {
	descendant := nd.left
	nd.left = descendant.right
	descendant.right = nd

	nd.fixHeight()
	descendant.fixHeight()

	return descendant
}

// rotateLeft performs the left
// rotation in order to balance the tree.
func (nd *node[TKey, TValue]) rotateLeft() *node[TKey, TValue] {
	descendant := nd.right
	nd.right = descendant.left
	descendant.left = nd

	nd.fixHeight()
	descendant.fixHeight()

	return descendant
}

// balance performs rotations in order
// to balance the AVL tree.
func (nd *node[TKey, TValue]) balance() *node[TKey, TValue] {
	nd.fixHeight()

	if nd.balanceFactor() == 2 {
		if nd.right.balanceFactor() < 0 {
			nd.right = nd.right.rotateRight()
		}

		return nd.rotateLeft()
	}

	if nd.balanceFactor() == -2 {
		if nd.left.balanceFactor() > 0 {
			nd.left = nd.rotateLeft()
		}

		return nd.rotateRight()
	}

	return nd
}

// search looks for the node with the
// specified key and returns its value.
//
// The second returned value indicates
// of the key exists in the AVL tree.
func (nd *node[TKey, TValue]) search(key TKey) (TValue, bool) {
	var zero TValue

	if nd == nil {
		return zero, false
	}

	if nd.key == key {
		return nd.value, true
	}

	if key < nd.key {
		if nd.left == nil {
			return zero, false
		}

		return nd.left.search(key)
	}

	if key > nd.key {
		if nd.right == nil {
			return zero, false
		}

		return nd.right.search(key)
	}

	return zero, false
}

// traverse traverses the tree performing
// the 'visit' function on each node. The
// order is depth-first.
func (nd *node[TKey, TValue]) traverse(visit func(currentKey TKey, currentValue TValue) error) error {
	if nd.left != nil {
		err := nd.left.traverse(visit)

		if err != nil {
			return err
		}
	}

	if nd != nil {
		err := visit(nd.key, nd.value)

		if err != nil {
			return err
		}
	}

	if nd.right != nil {
		err := nd.right.traverse(visit)

		if err != nil {
			return err
		}
	}

	return nil
}

// insert a value by the key in the
// tree (overwrites the previous value).
func (nd *node[TKey, TValue]) insert(key TKey, value TValue) *node[TKey, TValue] {
	if nd == nil {
		return &node[TKey, TValue]{
			key:   key,
			value: value,
		}
	}

	if key < nd.key {
		nd.left = nd.left.insert(key, value)
	} else {
		nd.right = nd.right.insert(key, value)
	}

	return nd.balance()
}

// findMin returns the node with
// the minimal value of the tree.
func (nd *node[TKey, TValue]) findMin() *node[TKey, TValue] {
	if nd.left != nil {
		nd.left.findMin()
	}

	return nd
}

// removeMin removes the node with
// the minimal value of the tree.
func (nd *node[TKey, TValue]) removeMin() *node[TKey, TValue] {
	if nd.left == nil {
		return nd.right
	}

	nd.left = nd.left.removeMin()

	return nd.balance()
}

// remove removes the node with
// the specified key from the tree.
func (nd *node[TKey, TValue]) remove(key TKey) *node[TKey, TValue] {
	if nd == nil {
		return nil
	}

	if key < nd.key {
		nd.left = nd.left.remove(key)
	} else if key > nd.key {
		nd.right = nd.right.remove(key)
	} else {
		left := nd.left
		right := nd.right

		min := right.findMin()
		min.right = right.removeMin()
		min.left = left

		return min.balance()
	}

	return nd.balance()
}
