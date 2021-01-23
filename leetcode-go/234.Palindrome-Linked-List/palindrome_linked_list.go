package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

type ListNode = structures.ListNode

func main() {
	arr := []int{1, 1, 2, 1}
	res := isPalindrome(structures.Ints2ListNode(arr))
	dfs(structures.Ints2ListNode(arr))
	fmt.Println(res)
}

/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

// @lc code=end

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
