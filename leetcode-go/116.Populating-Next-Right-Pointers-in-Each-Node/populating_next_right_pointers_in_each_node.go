package main

import (
	"github.com/onwl007/leetcode-algo/structures"
)

type Node = structures.Node

func connect(root *Node) *Node {
	// Time: O(N) Space: O(N)
	if root == nil {
		return nil
	}
	// 初始化队列，将根节点加入队列
	queue := []*Node{root}
	// 循环迭代层数，而不是每个节点
	for len(queue) != 0 {
		tmp := queue
		queue = nil
		// 遍历这一层的节点
		for i, node := range tmp {
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}
