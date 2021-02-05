package main

import "strings"

/*
 * @lc app=leetcode.cn id=331 lang=golang
 *
 * [331] 验证二叉树的前序序列化
 */

// @lc code=start
func isValidSerialization(preorder string) bool {
	slots := 1
	pros := strings.Split(preorder, ",")
	for _, v := range pros {
		slots--

		if slots < 0 {
			return false
		}

		if v != "#" {
			slots += 2
		}
	}
	return slots == 0
}

// @lc code=end

// 二叉树的前序遍历结果中，# 的个数正好是节点个数加 1，在未读完整个数组时，满足 # 的个数整好等于节点个数加 1
func isValidSerialization1(preorder string) bool {
	leaves, node := 0, 0
	pros := strings.Split(preorder, ",")
	for i, v := range pros {
		if v == "#" {
			leaves++
		} else {
			node++
		}
		// 未读完之前就出现大于，直接返回 false
		if leaves > node+1 {
			return false
		}
		// 在没有读到最后一个元素之前就满足了条件，直接返回 false
		if leaves == node+1 && i < len(pros)-1 {
			return false
		}
	}
	if leaves == node+1 {
		return true
	}
	return false
}
