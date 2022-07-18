package main

import (
	"fmt"
	"strconv"

	"github.com/zergon321/avl"
)

const (
	DataLength = 1000
)

func main() {
	tree := avl.NewTree[int, string]()

	for i := 0; i < DataLength; i++ {
		tree.Insert(i, strconv.Itoa(i))
	}

	tree.Traverse(func(currentKey int, currentValue string) error {
		fmt.Println(currentKey, currentValue)
		return nil
	})
}
