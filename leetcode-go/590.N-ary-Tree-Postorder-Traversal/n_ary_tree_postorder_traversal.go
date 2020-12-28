package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) (vals []int) {
	dfs(root, &vals)
	return
}

func dfs(root *Node, vals *[]int) {
	if root == nil {
		return
	}

	for _, v := range root.Children {
		dfs(v, vals)
	}
	*vals = append(*vals, root.Val)
}
