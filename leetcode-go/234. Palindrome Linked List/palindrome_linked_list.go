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

func dfs(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Println(head.Val)
	dfs(head.Next)
}
