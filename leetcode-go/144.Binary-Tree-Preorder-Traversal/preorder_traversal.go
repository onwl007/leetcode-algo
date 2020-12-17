package main

import (
	"github.com/onwl007/leetcode-algo/structures"
)

type TreeNode = structures.TreeNode

func preorderTraversal(root *TreeNode) (vals []int) {
	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		vals = append(vals, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}

func preorderTraversal1(root *TreeNode) (vals []int) {
	preorder(&vals, root)
	return
}

func preorder(vals *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	*vals = append(*vals, root.Val)
	preorder(vals, root.Left)
	preorder(vals, root.Right)
}

// 迭代法实现前序遍历二叉树
func preorderTraversalIteration(root *TreeNode) []int {
	vals := []int{}
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			vals = append(vals, root.Val)
			stack = append(stack, root.Right)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return vals
}

func preorderIteration(root *TreeNode) (vals []int) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		vals = append(vals, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return
}
