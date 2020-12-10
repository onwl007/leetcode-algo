package main

import "github.com/onwl007/leetcode-algo/structures"

type TreeNode = structures.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}
