package main

import (
	"fmt"
	"strconv"

	"github.com/onwl007/leetcode-algo/structures"
)

func main() {
	arr := []int{1, 2, 3, structures.NULL, 5}
	res := binaryTreePaths(structures.Ints2TreeNode(arr))
	fmt.Println(res)
}

type TreeNode = structures.TreeNode

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
func binaryTreePaths(root *TreeNode) []string {
	var vals []string
	var res string
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

// @lc code=end
