package main

import (
	"one-algorithm/internal"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	head := internal.BuildListNode([]int{1, 2, 3, 4, 5})
	node := reverseKGroup(head, 2)
	node.Print()
}
