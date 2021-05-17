package main

import (
	"fmt"
	"sort"

	"github.com/onwl007/leetcode-algo/structures"
)

type ListNode = structures.ListNode

/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	return sort1(head, nil)
}

func sort1(left, right *ListNode) *ListNode {
	if left == nil {
		return left
	}
	if left.Next == right {
		left.Next = nil
		return left
	}
	slow, fast := left, left
	for fast != right {
		slow = slow.Next
		fast = fast.Next
		if fast != right {
			fast = fast.Next
		}
	}
	mid := slow
	return mergeTwoList(sort1(left, mid), sort1(mid, right))
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
	prehead := &ListNode{}
	res := prehead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prehead.Next = l1
			l1 = l1.Next
		} else {
			prehead.Next = l2
			l2 = l2.Next
		}
		prehead = prehead.Next
	}
	if l1 != nil {
		prehead.Next = l1
	}
	if l2 != nil {
		prehead.Next = l2
	}
	return res.Next
}

// @lc code=end

func sortList1(head *ListNode) *ListNode {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	sort.Ints(arr)
	fmt.Println(arr)
	prehead := &ListNode{}
	res := prehead
	for _, v := range arr {
		prehead.Next = &ListNode{Val: v}
		prehead = prehead.Next
	}
	return res.Next
}
