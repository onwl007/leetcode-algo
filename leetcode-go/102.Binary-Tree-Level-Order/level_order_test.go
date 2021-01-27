package main

import (
	"testing"

	"github.com/onwl007/leetcode-algo/structures"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	in  *TreeNode
	out [][]int
}

func TestLevelOrder(t *testing.T) {
	qs := []Case{
		{in: structures.Ints2TreeNode([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}), out: [][]int{{3}, {9, 20}, {15, 7}}},
		{in: structures.Ints2TreeNode([]int{1, 2, 3, 4, 5, 6, 7}), out: [][]int{{1}, {2, 3}, {4, 5, 6, 7}}},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, levelOrder(v.in), "二叉树的层序遍历")
		ast.Equal(v.out, levelOrder1(v.in), "二叉树的层序遍历")
		ast.Equal(v.out, levelOrderRecursion(v.in), "二叉树的层序遍历")
	}
}
