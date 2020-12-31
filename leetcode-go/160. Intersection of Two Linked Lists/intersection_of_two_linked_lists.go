package main

import "github.com/onwl007/leetcode-algo/structures"

type ListNode = structures.ListNode

// map 法
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[interface{}]interface{})

	for headA != nil {
		m[headA] = 1
		headA = headA.Next
	}

	for headB != nil {
		if m[headB] == 1 {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

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
