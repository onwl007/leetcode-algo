package main

import (
	"math"

	"github.com/onwl007/leetcode-algo/structures"
)

type TreeNode = structures.TreeNode

func minDepth(root *TreeNode) int {
	// 其根节点为空时，其深度为 0
	if root == nil {
		return 0
	}
	// 判断是不是叶子节点，为叶子节点时返回 1，其深度为 1
	if root.Left == nil && root.Right == nil {
		return 1
	}
	v := math.MaxInt32
	// 递归其左子树，找出最小深度
	if root.Left != nil {
		v = min(minDepth(root.Left), v)
	}
	// 递归其右子树，找出最小深度
	if root.Right != nil {
		v = min(minDepth(root.Right), v)
	}
	// 深度加 1
	return v + 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
