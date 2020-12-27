package main

import (
	"github.com/onwl007/leetcode-algo/structures"
)

type TreeNode = structures.TreeNode

func flatten(root *TreeNode) {
	vals := []*TreeNode{}
	dfs(root, &vals)
	for i := 1; i < len(vals); i++ {
		prev, curr := vals[i-1], vals[i]
		prev.Left, prev.Right = nil, curr
	}
}

func dfs(root *TreeNode, vals *[]*TreeNode) {
	if root == nil {
		return
	}
	*vals = append(*vals, root)
	dfs(root.Left, vals)
	dfs(root.Right, vals)
}
