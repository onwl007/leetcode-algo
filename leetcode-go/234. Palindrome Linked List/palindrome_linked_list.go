package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

type ListNode = structures.ListNode

func isPalindrome(head *ListNode) bool {
	vals := []int{}
	for ; head != nil; head = head.Next {
		vals = append(vals, head.Val)
	}

	n := len(vals)
	for i, v := range vals[:n/2] {
		if v != vals[n-1-i] {
			return false
		}
	}
	return true
}

var left *ListNode

func isPalindrome1(head *ListNode) bool {
	left = head
	return traverse(head)
}

func traverse(right *ListNode) bool {
	if right == nil {
		return true
	}
	res := traverse(right.Next)
	res = res && (right.Val == left.Val)
	left = left.Next
	return res
}

func dfs(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Println(head.Val)
	dfs(head.Next)
}
