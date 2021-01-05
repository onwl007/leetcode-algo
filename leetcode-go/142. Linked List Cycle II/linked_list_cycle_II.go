package main

import "github.com/onwl007/leetcode-algo/structures"

type ListNode = structures.ListNode

// 哈希解法
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return nil
}

// 快慢指针解法
func detectCycle1(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			slow = head
			for slow != fast {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
