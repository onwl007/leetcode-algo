package main

import "github.com/onwl007/leetcode-algo/structures"

type TreeNode = structures.TreeNode

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBST(root, nil, nil)
}

func isBST(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	return isBST(root.Left, min, root) && isBST(root.Right, root, max)
}
