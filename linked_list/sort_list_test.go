package main

import (
	"one-algorithm/internal"
	"testing"
)

func TestSortList(t *testing.T) {
	head := internal.BuildListNode([]int{4, 2, 1, 3})
	node := sortList(head)
	node.Print()
}
