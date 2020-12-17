package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

func main() {
	// arr := []int{1, structures.NULL, 2, 3}
	arr := []int{1, 2, 3, structures.NULL, 4, 5, 6}
	vals := preorderTraversalIteration(structures.Ints2TreeNode(arr))
	fmt.Println(vals)
}
