package main

import "github.com/onwl007/leetcode-algo/structures"

type ListNode = structures.ListNode

func middleNode(head *ListNode) *ListNode {
	vals := []*ListNode{}
	for head != nil {
		vals = append(vals, head)
		head = head.Next
	}
	return vals[len(vals)/2]
}

// 两次遍历
func middleNode1(head *ListNode) *ListNode {
	n := 0
	cur := head
	for head != nil {
		n++
		head = head.Next
	}

	for i := 0; i < n/2; i++ {
		cur = cur.Next
	}
	return cur
}

// 快慢指针
func middleNode2(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
