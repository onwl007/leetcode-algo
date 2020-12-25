package main

import (
	"strconv"

	"github.com/onwl007/leetcode-algo/structures"
)

type TreeNode = structures.TreeNode

func binaryTreePaths(root *TreeNode) []string {
	vals := []string{}
	res := ""
	dfs(root, &vals, res)
	return vals
}

func dfs(root *TreeNode, vals *[]string, res string) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		res += strconv.Itoa(root.Val)
		*vals = append(*vals, res)
		return
	}
	res += strconv.Itoa(root.Val) + "->"
	dfs(root.Left, vals, res)
	dfs(root.Right, vals, res)
}
