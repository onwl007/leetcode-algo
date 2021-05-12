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

// @lc code=end
