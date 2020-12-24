package main

import "github.com/onwl007/leetcode-algo/structures"

type TreeNode = structures.TreeNode

// 二叉搜索树的中序遍历是有序的，所以数组的第 k-1 个元素就是第 k 小的元素
func kthSmallest(root *TreeNode, k int) int {
	vals := []int{}
	inorder(&vals, root)
	return vals[k-1]
}

func inorder(vals *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	inorder(vals, root.Left)
	*vals = append(*vals, root.Val)
	inorder(vals, root.Right)
}
