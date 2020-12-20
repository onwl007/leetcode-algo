package main

import "github.com/onwl007/leetcode-algo/structures"

type TreeNode = structures.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	val := preorder[0]
	root := &TreeNode{Val: val}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == val {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:1])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
