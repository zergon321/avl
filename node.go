package avl

import (
	"golang.org/x/exp/constraints"
)

type node[TKey constraints.Ordered, TValue any] struct {
	height int
	left   *node[TKey, TValue]
	right  *node[TKey, TValue]
	key    TKey
	value  TValue
}

func (nd *node[TKey, TValue]) getHeight() int {
	if nd == nil {
		return 0
	}

	return nd.height
}

func (nd *node[TKey, TValue]) balanceFactor() int {
	return nd.right.getHeight() - nd.left.getHeight()
}

func (nd *node[TKey, TValue]) fixHeight() {
	leftHeight := nd.left.getHeight()
	rightHeight := nd.right.getHeight()

	if leftHeight > rightHeight {
		nd.height = leftHeight + 1
	} else {
		nd.height = rightHeight + 1
	}
}

func (nd *node[TKey, TValue]) rotateRight() *node[TKey, TValue] {
	descendant := nd.left
	nd.left = descendant.right
	descendant.right = nd

	nd.fixHeight()
	descendant.fixHeight()

	return descendant
}

func (nd *node[TKey, TValue]) rotateLeft() *node[TKey, TValue] {
	descendant := nd.right
	nd.right = descendant.left
	descendant.left = nd

	nd.fixHeight()
	descendant.fixHeight()

	return descendant
}

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

func (nd *node[TKey, TValue]) findMin() *node[TKey, TValue] {
	if nd.left != nil {
		nd.left.findMin()
	}

	return nd
}

func (p *node[TKey, TValue]) removeMin() *node[TKey, TValue] {
	if p.left == nil {
		return p.right
	}

	p.left = p.left.removeMin()

	return p.balance()
}

func (p *node[TKey, TValue]) remove(key TKey) *node[TKey, TValue] {
	if p == nil {
		return nil
	}

	if key < p.key {
		p.left = p.left.remove(key)
	} else if key > p.key {
		p.right = p.right.remove(key)
	} else {
		q := p.left
		r := p.right

		min := r.findMin()
		min.right = r.removeMin()
		min.left = q

		return min.balance()
	}

	return p.balance()
}
