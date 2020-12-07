package main

import "github.com/onwl007/leetcode-algo/structures"

type TreeNode = structures.TreeNode

func inorderTraversal(root *TreeNode) (vals []int) {
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		vals = append(vals, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return vals
}
