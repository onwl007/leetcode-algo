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
