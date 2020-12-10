package main

import "github.com/onwl007/leetcode-algo/structures"

type TreeNode = structures.TreeNode

// 构成对称树的条件
// 1. 根节点相同
// 2. 每个树的左子树都与另外一个树的右子树相等
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameTree(root, root)
}

// 利用判断两个树相同的方法
// 递归时只需把一个树的左子树与另一个的右子树相比较
// 一个树的右子树与另一个的左子树相比较
func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Right) && isSameTree(p.Right, q.Left)
}
