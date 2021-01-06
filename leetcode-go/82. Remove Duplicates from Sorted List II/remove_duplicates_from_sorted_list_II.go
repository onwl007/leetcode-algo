package main

import (
	"sort"

	"github.com/onwl007/leetcode-algo/structures"
)

type ListNode = structures.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	m := make(map[int]int)
	for head != nil {
		if _, ok := m[head.Val]; ok {
			m[head.Val]++
		} else {
			m[head.Val] = 1
		}
		head = head.Next
	}

	vals := []int{}
	for k, v := range m {
		if v == 1 {
			vals = append(vals, k)
		}
	}
	sort.Ints(vals)

	dummy := &ListNode{}
	pre := dummy
	for _, v := range vals {
		tmp := &ListNode{Val: v}
		pre.Next = tmp
		pre = pre.Next
	}
	return dummy.Next
}
