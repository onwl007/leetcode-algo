package main

import (
	"testing"

	"github.com/onwl007/leetcode-algo/structures"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	head *ListNode
	out  []int
}

func TestReverseList(t *testing.T) {
	qs := []Case{
		{head: structures.Ints2ListNode([]int{1, 2, 3, 4, 5}), out: []int{5, 4, 3, 2, 1}},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, structures.ListNode2Ints(reverseList(v.head)), "反转链表")
		// ast.Equal(v.out, structures.ListNode2Ints(reverseList1(v.head)), "反转链表")
	}
}
