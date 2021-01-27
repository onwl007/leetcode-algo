package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

type TreeNode = structures.TreeNode

func main() {
	arr := []int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}
	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	// 一维数组返回
	vals1 := levelOrder1(structures.Ints2TreeNode(arr1))
	fmt.Println(vals1)
	fmt.Println("-----------")
	// 二维数组返回
	vals2 := levelOrder(structures.Ints2TreeNode(arr))
	fmt.Println(vals2)
}

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	vals := [][]int{}
	if root == nil {
		return vals
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		size := len(queue)
		res := []int{}
		for i := 0; i < size; i++ {
			node := queue[i]
			res = append(res, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		vals = append(vals, res)
		queue = queue[size:]
		res = []int{}
	}
	return vals
}

// @lc code=end

func levelOrder1(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

func levelOrder2(root *TreeNode) (vals []int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		node := queue[0]
		vals = append(vals, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		queue = queue[1:]
	}
	return
}
