package leetcode

import (
	"github.com/onwl007/leetcode-algo/structures"
)

type ListNode = structures.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tmp := 0
	list := &ListNode{Val: 0}
	head := list

	for l1 != nil || l2 != nil || tmp != 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}

		list.Next = &ListNode{Val: 0}
		tmp = tmp / 10
		list = list.Next
	}
	return head.Next
}
