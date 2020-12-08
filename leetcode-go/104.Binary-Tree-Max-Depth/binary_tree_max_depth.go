package main

import "github.com/onwl007/leetcode-algo/structures"

type TreeNode = structures.TreeNode

// 只需要求左子树和右子树的最大深度 l 和 r，那么该二叉树的最大深度即为 max(l,r) + 1
func maxDepth(root *TreeNode) int {
	// 递归终止条件
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
