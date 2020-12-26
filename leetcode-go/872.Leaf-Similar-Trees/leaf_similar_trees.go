package main

import (
	"reflect"

	"github.com/onwl007/leetcode-algo/structures"
)

type TreeNode = structures.TreeNode

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	vals1 := []int{}
	vals2 := []int{}
	dfs(root1, &vals1)
	dfs(root2, &vals2)
	return reflect.DeepEqual(vals1, vals2)
}

func dfs(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	// 判断是叶子节点
	if root.Left == nil && root.Right == nil {
		*vals = append(*vals, root.Val)
	}
	dfs(root.Left, vals)
	dfs(root.Right, vals)
}
