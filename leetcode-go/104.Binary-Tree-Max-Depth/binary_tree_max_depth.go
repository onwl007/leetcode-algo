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

func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]   // 取队头元素
			queue = queue[1:] // 队头元素出队
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		depth++
	}
	return depth
}
