package main

import "github.com/onwl007/leetcode-algo/structures"

type TreeNode = structures.TreeNode

// 递归解法
// 二叉搜索树，根节点大于左子树所有的值，小于右子树所有的值
// 当 p，q 同时小于根节点，递归左子树
// 当 p，q 同时大于根节点，递归右子树
func lowestCommonAncestorRecurison(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestorRecurison(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestorRecurison(root.Right, p, q)
	}
	return root
}

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val < root.Val && q.Val < root.Val {
			root = root.Right
		} else {
			break
		}
	}
	return root
}
