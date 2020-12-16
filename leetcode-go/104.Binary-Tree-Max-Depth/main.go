package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

func main() {
	arr := []int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}
	depth := maxDepthBFS(structures.Ints2TreeNode(arr))
	fmt.Println(depth)
}
