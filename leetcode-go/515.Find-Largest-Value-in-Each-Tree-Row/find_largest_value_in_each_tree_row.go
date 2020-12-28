package main

import "github.com/onwl007/leetcode-algo/structures"

type TreeNode = structures.TreeNode

func largestValues(root *TreeNode) []int {
	vals := []int{}

	if root == nil {
		return vals
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		tmp := queue
		queue = nil
		max := -1 << 63
		for _, v := range tmp {
			if v.Val > max {
				max = v.Val
			}
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		vals = append(vals, max)
	}
	return vals
}

func largestValues1(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	vals := []int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		max := -1 << 63
		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		vals = append(vals, max)
		queue = queue[size:]
	}
	return vals
}

func largestValuesRecurison(root *TreeNode) []int {
	vals := []int{}
	dfs(root, &vals, 0)
	return vals
}

func dfs(root *TreeNode, vals *[]int, level int) {
	if root == nil {
		return
	}

	if len(*vals) == level {
		*vals = append(*vals, root.Val)
	}
	(*vals)[level] = max((*vals)[level], root.Val)
	dfs(root.Left, vals, level+1)
	dfs(root.Right, vals, level+1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
