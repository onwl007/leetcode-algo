package main

import "github.com/onwl007/leetcode-algo/structures"

func main() {
	arr := []int{1, 2, 5, 3, 4, structures.NULL, 6}
	flatten(structures.Ints2TreeNode(arr))
}
