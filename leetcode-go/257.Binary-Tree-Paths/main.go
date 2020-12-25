package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

func main() {
	arr := []int{1, 2, 3, structures.NULL, 5}
	res := binaryTreePaths(structures.Ints2TreeNode(arr))
	fmt.Println(res)
}
