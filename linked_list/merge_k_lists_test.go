package main

import (
	. "one-algorithm/internal"
	"testing"
)

func TestMergeKLists(t *testing.T) {

	lists := []*ListNode{BuildListNode([]int{1, 4, 5}),
		BuildListNode([]int{1, 3, 4}),
		BuildListNode([]int{2, 6})}
	node := mergeKLists(lists)
	node.Print()
}
