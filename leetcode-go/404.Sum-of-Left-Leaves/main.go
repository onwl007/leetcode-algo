package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

func main() {
	arr := []int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}
	leftSum := sumOfLeftLeaves(structures.Ints2TreeNode(arr))
	rightSum := sumOfRightLeaves(structures.Ints2TreeNode(arr))
	fmt.Printf("sum of left leaves ----- %v\n", leftSum)
	fmt.Printf("sum of right leaves ----- %v\n", rightSum)
}
