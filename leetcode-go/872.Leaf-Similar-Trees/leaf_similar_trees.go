package main

import (
	"fmt"
	"reflect"

	"github.com/onwl007/leetcode-algo/structures"
)

type TreeNode = structures.TreeNode

func main() {
	arr1 := []int{3, 5, 1, 6, 2, 9, 8, structures.NULL, structures.NULL, 7, 4}
	arr2 := []int{3, 5, 1, 6, 7, 4, 2, structures.NULL, structures.NULL, structures.NULL, structures.NULL, structures.NULL, structures.NULL, 9, 8}
	t := leafSimilar(structures.Ints2TreeNode(arr1), structures.Ints2TreeNode(arr2))
	fmt.Println(t)
}

/*
 * @lc app=leetcode.cn id=872 lang=golang
 *
 * [872] 叶子相似的树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
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

// @lc code=end
