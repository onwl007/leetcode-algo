package main

import (
	"testing"

	"github.com/onwl007/leetcode-algo/structures"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	root *TreeNode
	vals []int
}

func TestLargestValues(t *testing.T) {
	qs := []Case{
		{root: structures.Ints2TreeNode([]int{1, 3, 2, 5, 3, structures.NULL, 9}), vals: []int{1, 3, 9}},
		{root: structures.Ints2TreeNode([]int{}), vals: []int{}},
		{root: structures.Ints2TreeNode([]int{1}), vals: []int{1}},
		{root: structures.Ints2TreeNode([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}), vals: []int{3, 20, 15}},
	}

	ast := assert.New(t)

	for _, v := range qs {
		ast.Equal(v.vals, largestValues(v.root), "在每个树行中找最大值")
		ast.Equal(v.vals, largestValues1(v.root), "在每个树行中找最大值")
		ast.Equal(v.vals, largestValuesRecurison(v.root), "在每个树行中找最大值")
	}
}
