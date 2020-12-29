package structures

import (
	"fmt"
)

// ListNode 是链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// List2Ints convert List to []int
func ListNode2Ints(head *ListNode) []int {
	limit := 100
	times := 0
	res := []int{}

	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// Ints2List convert []int to List
func Ints2ListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}
