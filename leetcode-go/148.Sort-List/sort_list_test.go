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

func TestSortList(t *testing.T) {
	node1 := structures.Ints2ListNode([]int{4, 2, 1, 3})
	node2 := structures.Ints2ListNode([]int{-1, 5, 3, 4, 0})
	qs := []Case{
		{head: node1, out: []int{1, 2, 3, 4}},
		{head: node2, out: []int{-1, 0, 3, 4, 5}},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.ElementsMatch(v.out, structures.ListNode2Ints(sortList(v.head)), "148 排序链表")
	}
}
