package main

import (
	"testing"

	"github.com/onwl007/leetcode-algo/structures"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	lists []*ListNode
	out   []int
}

func TestMergeKLists(t *testing.T) {
	node1 := structures.Ints2ListNode([]int{1, 4, 5})
	node2 := structures.Ints2ListNode([]int{1, 3, 4})
	node3 := structures.Ints2ListNode([]int{2, 6})
	qs := []Case{
		{lists: []*ListNode{node1, node2, node3}, out: []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{lists: []*ListNode{}, out: []int{}},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.ElementsMatch(v.out, structures.ListNode2Ints(mergeKLists(v.lists)), "合并 K 个升序链表")
		// 	ast.ElementsMatch(v.out, structures.ListNode2Ints(mergeKLists1(v.lists)), "合并 K 个升序链表")
	}
}
