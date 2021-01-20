package main

import (
	"testing"

	"github.com/onwl007/leetcode-algo/structures"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	in  *TreeNode
	out bool
}

func TestPathSum(t *testing.T) {
	qs := []Case{
		{in: structures.Ints2TreeNode([]int{5, 4, 8, 11, 13, 4, 7, 2, structures.NULL, structures.NULL, structures.NULL, 1}), out: true},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, hasPathSum(v.in, 22), "路径总和")
	}
}
