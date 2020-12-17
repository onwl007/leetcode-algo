package main

import "github.com/onwl007/leetcode-algo/structures"

type TreeNode = structures.TreeNode

func postorderTraversal(root *TreeNode) (vals []int) {
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		vals = append(vals, node.Val)
	}
	postorder(root)
	return
}

func postorderTraversal1(root *TreeNode) (vals []int) {
	postorder(&vals, root)
	return
}

func postorder(vals *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	postorder(vals, root.Left)
	postorder(vals, root.Right)
	*vals = append(*vals, root.Val)
}
