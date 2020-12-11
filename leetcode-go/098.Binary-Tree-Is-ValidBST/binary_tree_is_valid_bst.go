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
	// 判断右子树的值小于其根节点的值，不符合 二叉搜索树
	if min != nil && root.Val <= min.Val {
		return false
	}
	// 判断左子树的节点大于其根节点的值，不符合二叉搜索树
	if max != nil && root.Val >= max.Val {
		return false
	}
	return isBST(root.Left, min, root) && isBST(root.Right, root, max)
}
