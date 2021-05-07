package main

import "github.com/onwl007/leetcode-algo/structures"

type ListNode = structures.ListNode

func main() {
	l1 := structures.Ints2ListNode([]int{1, 2, 4})
	l2 := structures.Ints2ListNode([]int{1, 3, 4})
	mergeTwoLists(l1, l2)
}

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{}
	res := prehead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prehead.Next = l1
			l1 = l1.Next
		} else {
			prehead.Next = l2
			l2 = l2.Next
		}
		prehead = prehead.Next
	}

	if l1 != nil {
		prehead.Next = l1
	}
	if l2 != nil {
		prehead.Next = l2
	}
	return res.Next
}

// @lc code=end

// 递归解法
func mergeTwoLists1(l1, l2 *ListNode) *ListNode  {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists1(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists1(l1, l2.Next)
		return l2
	}
}
