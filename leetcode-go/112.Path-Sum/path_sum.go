package main

import "github.com/onwl007/leetcode-algo/structures"

type TreeNode = structures.TreeNode

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	// 递归左子树或右子树，转换为 sum 减去当前节点的值，到达叶子节点时，sum 的值应正好等于 0
	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
