package main

import (
	. "one-algorithm/pkg"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	head := BuildListNode([]int{1, 2, 3, 4, 5})
	reverseBetween(head, 2, 4)
	head.Print()
}
