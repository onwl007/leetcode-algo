package main

import "github.com/onwl007/leetcode-algo/structures"

type ListNode = structures.ListNode

// map
/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]int)

	for headA != nil {
		m[headA]++
		headA = headA.Next
	}

	for headB != nil {
		if _, ok := m[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// @lc code=end

// 暴力法
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	curA := headA
	for curA != nil {
		curB := headB
		for curB != nil {
			if curA == curB {
				return curA
			}
			curB = curB.Next
		}
		curA = curA.Next
	}
	return nil
}
