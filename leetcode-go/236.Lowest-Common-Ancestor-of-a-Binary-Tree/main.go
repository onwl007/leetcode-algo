package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

func main() {
	root := structures.Ints2TreeNode([]int{3, 5, 1, 6, 2, 0, 8, structures.NULL, structures.NULL, 7, 4})
	p := structures.Ints2TreeNode([]int{7})
	q := structures.Ints2TreeNode([]int{4})
	lca := lowestCommonAncestor(root, p, q)
	fmt.Printf("最近公共祖先节点 ---- %v \n", lca.Val)
}
