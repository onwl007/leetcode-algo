package main

import (
	"testing"

	"github.com/onwl007/leetcode-algo/structures"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	head *ListNode
	val  int
	out  []int
}

func TestRemoveElements(t *testing.T) {
	qs := []Case{
		{head: structures.Ints2ListNode([]int{1, 2, 6, 3, 4, 5, 6}), val: 6, out: []int{1, 2, 3, 4, 5}},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.ElementsMatch(v.out, structures.ListNode2Ints(removeElements(v.head, v.val)), "移除链表元素")
	}
}
