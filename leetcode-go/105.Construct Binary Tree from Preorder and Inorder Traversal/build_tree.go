package main

import "github.com/onwl007/leetcode-algo/structures"

type TreeNode = structures.TreeNode

// 题解
// 前序遍历第一个元素的值是根节点，拿到根节点的值，可以找到根节点在中序遍历的位置，从而知道左子树和右子树，
// 此时，再去递归左子树和右子树即可构造出一颗二叉树
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
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])   // 前序遍历取根节点后面到左子树的长度加 1 的区间
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:]) // 左子树的长度加 1 后面的区间，也就是右子树
	return root
}
