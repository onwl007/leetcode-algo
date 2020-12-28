package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) (vals []int) {
	dfs(root, &vals)
	return
}

func dfs(root *Node, vals *[]int) {
	if root == nil {
		return
	}

	*vals = append(*vals, root.Val)
	for _, v := range root.Children {
		dfs(v, vals)
	}
}
