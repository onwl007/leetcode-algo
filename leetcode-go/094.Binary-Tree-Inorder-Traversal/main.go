package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, structures.NULL}
	vals := inorderIteration(structures.Ints2TreeNode(arr))
	fmt.Println(vals)
}
