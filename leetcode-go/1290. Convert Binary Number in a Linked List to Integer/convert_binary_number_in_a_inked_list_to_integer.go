package main

import "github.com/onwl007/leetcode-algo/structures"

type ListNode = structures.ListNode

func getDecimalValue(head *ListNode) int {
	sum := 0
	for head != nil {
		sum = (sum << 1) + head.Val
		head = head.Next
	}
	return sum
}
