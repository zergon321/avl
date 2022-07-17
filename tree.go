package avl

import "golang.org/x/exp/constraints"

type Tree[TKey constraints.Ordered, TValue any] struct {
	root *node[TKey, TValue]
}
