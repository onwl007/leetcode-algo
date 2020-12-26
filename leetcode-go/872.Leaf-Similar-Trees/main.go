package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

func main() {
	arr1 := []int{3, 5, 1, 6, 2, 9, 8, structures.NULL, structures.NULL, 7, 4}
	arr2 := []int{3, 5, 1, 6, 7, 4, 2, structures.NULL, structures.NULL, structures.NULL, structures.NULL, structures.NULL, structures.NULL, 9, 8}
	t := leafSimilar(structures.Ints2TreeNode(arr1), structures.Ints2TreeNode(arr2))
	fmt.Println(t)
}
