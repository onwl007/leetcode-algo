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
	// golang 每次遍历 map 的顺序是不一样的
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

func deleteDuplicates1(head *ListNode) *ListNode {
	dumy := &ListNode{}
	dumy.Next = head
	a := dumy
	b := head
	for b != nil && b.Next != nil {
		if a.Next.Val != b.Next.Val {
			a = a.Next
			b = b.Next
		} else {
			for b != nil && b.Next != nil && a.Next.Val == b.Next.Val {
				b = b.Next
			}
			a.Next = b.Next
			b = b.Next
		}
	}
	return dumy.Next
}
