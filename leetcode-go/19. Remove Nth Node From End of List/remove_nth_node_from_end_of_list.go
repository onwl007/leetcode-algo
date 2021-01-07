package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

type ListNode = structures.ListNode

// 循环遍历太多，可以减少循环
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for n > 0 {
		fast = fast.Next
		n--
	}

	k := 0
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
		k++
	}
	dummy := &ListNode{Val: 0, Next: head}
	cur := dummy
	for i := 0; i < k; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

// 对链表进行操作时，一种常用的技巧是添加一个哑结点
// 删除节点时，我们要知道它的前驱节点，将前驱节点执行其后继结点，这样就实现了删除节点
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	length := getLength(head)
	fmt.Println(length)
	dummy := &ListNode{Val: 0, Next: head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}
