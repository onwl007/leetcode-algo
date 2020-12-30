package main

import "github.com/onwl007/leetcode-algo/structures"

type ListNode = structures.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	h := &ListNode{Val: 0, Next: head}
	pre := h
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return h.Next
}
