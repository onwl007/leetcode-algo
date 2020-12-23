package main

import "github.com/onwl007/leetcode-algo/structures"

type TreeNode = structures.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p == root || q == root {
		return nil
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil && right == nil {
		return nil
	}
	if left == nil {
		return right
	}
	return left
}
