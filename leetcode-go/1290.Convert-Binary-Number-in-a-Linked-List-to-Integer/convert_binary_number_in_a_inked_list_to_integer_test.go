package main

import (
	"testing"

	"github.com/onwl007/leetcode-algo/structures"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	head *ListNode
	out  int
}

func TestGetDecimalValue(t *testing.T) {
	qs := []Case{
		{head: structures.Ints2ListNode([]int{1, 0, 1}), out: 5},
		{head: structures.Ints2ListNode([]int{0}), out: 0},
		{head: structures.Ints2ListNode([]int{1}), out: 1},
		{head: structures.Ints2ListNode([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}), out: 18880},
		{head: structures.Ints2ListNode([]int{0, 0}), out: 0},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, getDecimalValue(v.head), "二进制链表转整数")
	}
}
