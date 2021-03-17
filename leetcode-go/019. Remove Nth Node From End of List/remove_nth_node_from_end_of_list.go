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

// 可以通过栈的这种形式，遍历的时候一次入栈
// 我们弹出栈的第 nn 个节点就是需要删除的节点，
// 并且目前栈顶的节点就是待删除节点的前驱节点
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	vals := []*ListNode{}
	dummy := &ListNode{Val: 0, Next: head}
	for node := dummy; node != nil; node = node.Next {
		vals = append(vals, node)
	}
	pre := vals[len(vals)-n-1]
	pre.Next = pre.Next.Next
	return dummy.Next
}

// 快慢指针
// 快慢指针同时指向 head 结点时，快指针先移动 n 次，
// 然后快慢指针同时移动，当快指针移动到链表尾部时，慢指针此时正好移动到倒数第 n 个结点
// 本题需要删除倒数第 n 个节点，我们把慢指针指向 dummy 哑结点，其余步骤不变，
// 此时快指针移动链表尾部，慢指针正好移动到倒数第 n 个节点的前驱结点
// 我们正好把前驱结点指向其后继结点，这样就实现了删除倒数第 n 个节点
func removeNthFromEnd3(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	slow, fast := dummy, head
	for n > 0 {
		fast = fast.Next
		n--
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
