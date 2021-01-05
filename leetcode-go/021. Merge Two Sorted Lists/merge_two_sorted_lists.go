package main

import "github.com/onwl007/leetcode-algo/structures"

type ListNode = structures.ListNode

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
