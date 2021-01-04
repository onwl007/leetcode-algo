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
