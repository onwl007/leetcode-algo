package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

type TreeNode = structures.TreeNode

func main() {
	arr := []int{1, 3, 2, 5, 3, structures.NULL, 9}
	vals1 := largestValues(structures.Ints2TreeNode(arr))
	vals2 := largestValues1(structures.Ints2TreeNode(arr))
	vals3 := largestValuesRecurison(structures.Ints2TreeNode(arr))
	fmt.Println(vals1)
	fmt.Println(vals2)
	fmt.Println(vals3)
}

/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
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
func largestValues(root *TreeNode) []int {
	vals := []int{}

	if root == nil {
		return vals
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		tmp := queue
		queue = nil
		max := -1 << 63
		for _, v := range tmp {
			if v.Val > max {
				max = v.Val
			}
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		vals = append(vals, max)
	}
	return vals
}

// @lc code=end

func largestValues1(root *TreeNode) []int {
	vals := []int{}
	if root == nil {
		return vals
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		max := -1 << 63
		for i := 0; i < size; i++ {
			node := queue[i]
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		vals = append(vals, max)
		queue = queue[size:]
	}
	return vals
}

func largestValuesRecurison(root *TreeNode) []int {
	vals := []int{}
	dfs(root, &vals, 0)
	return vals
}

func dfs(root *TreeNode, vals *[]int, level int) {
	if root == nil {
		return
	}

	if len(*vals) == level {
		*vals = append(*vals, root.Val)
	}
	(*vals)[level] = max((*vals)[level], root.Val)
	dfs(root.Left, vals, level+1)
	dfs(root.Right, vals, level+1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
