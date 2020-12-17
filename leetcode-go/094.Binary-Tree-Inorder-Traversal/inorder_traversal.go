package main

import "github.com/onwl007/leetcode-algo/structures"

type TreeNode = structures.TreeNode

func inorderTraversal(root *TreeNode) (vals []int) {
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		vals = append(vals, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return vals
}

func inorderTraversal1(root *TreeNode) (vals []int) {
	inorder(&vals, root)
	return
}

func inorder(vals *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	inorder(vals, root.Left)
	*vals = append(*vals, root.Val)
	inorder(vals, root.Right)
}

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
