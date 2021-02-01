package main

import (
	"testing"

	"github.com/onwl007/leetcode-algo/structures"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	p   *TreeNode
	q   *TreeNode
	out bool
}

func TestIsSameTree(t *testing.T) {
	qs := []Case{
		{p: structures.Ints2TreeNode([]int{1, 2, 3}), q: structures.Ints2TreeNode([]int{1, 2, 3}), out: true},
		{p: structures.Ints2TreeNode([]int{1, 2}), q: structures.Ints2TreeNode([]int{1, structures.NULL, 2}), out: false},
		{p: structures.Ints2TreeNode([]int{1, 2, 1}), q: structures.Ints2TreeNode([]int{1, 1, 2}), out: false},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, isSameTree(v.p, v.q), "相同的树")
	}
}
