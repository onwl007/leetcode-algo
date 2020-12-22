package main

import (
	"strconv"
	"strings"

	"github.com/onwl007/leetcode-algo/structures"
)

type TreeNode = structures.TreeNode

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return prorder(root, "")
}

func prorder(root *TreeNode, str string) string {
	if root == nil {
		str += "null,"
	} else {
		str += strconv.Itoa(root.Val) + ","
		str = prorder(root.Left, str)
		prorder(root.Right, str)
	}
	return str
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	return buildTree(nodes)
}

func buildTree(nodes []string) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}
	val := nodes[0]
	nodes = nodes[1:]
	v, _ := strconv.Atoi(val)
	root := &TreeNode{Val: v}
	root.Left = buildTree(nodes)
	root.Right = buildTree(nodes)
	return root
}
