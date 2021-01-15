package main

import "github.com/onwl007/leetcode-algo/structures"

type ListNode = structures.ListNode

// 这道题其实不是直接删除节点 node
// 而是用 node 的下一个节点覆盖 node 节点
// 真正删除的其实是 node 的下一个节点

/*
 * @lc app=leetcode.cn id=237 lang=golang
 *
 * [237] 删除链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// @lc code=end
