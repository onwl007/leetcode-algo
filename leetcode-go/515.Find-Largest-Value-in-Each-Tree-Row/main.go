package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

func main() {
	arr := []int{1, 3, 2, 5, 3, structures.NULL, 9}
	vals1 := largestValues(structures.Ints2TreeNode(arr))
	vals2 := largestValues1(structures.Ints2TreeNode(arr))
	vals3 := largestValuesRecurison(structures.Ints2TreeNode(arr))
	fmt.Println(vals1)
	fmt.Println(vals2)
	fmt.Println(vals3)
}
