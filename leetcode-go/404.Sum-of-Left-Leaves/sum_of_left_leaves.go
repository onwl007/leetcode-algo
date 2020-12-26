package main

import "github.com/onwl007/leetcode-algo/structures"

type TreeNode = structures.TreeNode

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	dfs(root, &sum)
	return sum
}

func dfs(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	// 判断节点的左节点不为空并且，左节点的左节点和左节点的右节点同时为空
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		*sum += root.Left.Val
	}
	dfs(root.Left, sum)
	dfs(root.Right, sum)
}

func sumOfRightLeaves(root *TreeNode) int {
	sum := 0
	dfs1(root, &sum)
	return sum
}

func dfs1(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	// 右叶子
	// 节点的右节点不为空，右节点的左节点和右节点同时为空
	if root.Right != nil && root.Right.Left == nil && root.Right.Right == nil {
		*sum += root.Right.Val
	}
	dfs1(root.Left, sum)
	dfs1(root.Right, sum)
}
