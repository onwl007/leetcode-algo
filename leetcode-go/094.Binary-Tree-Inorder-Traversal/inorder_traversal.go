package main

import "github.com/onwl007/leetcode-algo/structures"

type TreeNode = structures.TreeNode

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	vals := []int{}
	dfs(root, &vals)
	return vals
}

func dfs(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, vals)
	*vals = append(*vals, root.Val)
	dfs(root.Right, vals)
}

// @lc code=end

func inorderIteration(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		vals = append(vals, root.Val)
		root = root.Right
	}
	return
}
