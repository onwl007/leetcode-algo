package main

import "github.com/onwl007/leetcode-algo/structures"

type ListNode = structures.ListNode

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
