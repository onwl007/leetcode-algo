package main

import (
	"github.com/onwl007/leetcode-algo/structures"
)

type TreeNode = structures.TreeNode

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	dfs(&res, root, sum, []int{})
	return res
}

func dfs(res *[][]int, root *TreeNode, sum int, path []int) {
	if root == nil {
		return
	}
	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			r := make([]int, len(path))
			// 将 path 数组拷贝到每次都声明的新数组中，如果不拷贝，path 在每次遍历中都会被改变
			copy(r, path)
			*res = append(*res, r)
			return
		}
	}
	dfs(res, root.Left, sum-root.Val, path)
	dfs(res, root.Right, sum-root.Val, path)
}
