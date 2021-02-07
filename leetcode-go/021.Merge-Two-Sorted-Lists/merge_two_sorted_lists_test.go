package main

import (
	"testing"

	"github.com/onwl007/leetcode-algo/structures"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	l1  *ListNode
	l2  *ListNode
	out []int
}

func TestMergeTwoLists(t *testing.T) {
	qs := []Case{
		{l1: structures.Ints2ListNode([]int{1, 2, 4}), l2: structures.Ints2ListNode([]int{1, 3, 4}), out: []int{1, 1, 2, 3, 4, 4}},
		{l1: structures.Ints2ListNode([]int{}), l2: structures.Ints2ListNode([]int{}), out: []int{}},
		{l1: structures.Ints2ListNode([]int{}), l2: structures.Ints2ListNode([]int{0}), out: []int{0}},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.ElementsMatch(v.out, structures.ListNode2Ints(mergeTwoLists(v.l1, v.l2)), "合并两个有序链表")
	}
}
