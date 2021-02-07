package main

import "github.com/onwl007/leetcode-algo/structures"

type ListNode = structures.ListNode

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 归并合并
func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	first := merge(lists, left, mid)
	second := merge(lists, mid+1, right)
	return mergeTwoLists(first, second)
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
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

// 顺序合并
func mergeKLists1(lists []*ListNode) *ListNode {
	prehead := &ListNode{}
	for i := 0; i < len(lists); i++ {
		prehead.Next = mergeTwoLists(prehead.Next, lists[i])
	}
	return prehead.Next
}
