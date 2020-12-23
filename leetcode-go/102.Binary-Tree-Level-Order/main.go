package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

func main() {
	arr := []int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}
	arr1 := []int{1, 2, 3, 4, 5, 6, 7}
	// 一维数组返回
	vals1 := levelOrder1(structures.Ints2TreeNode(arr1))
	fmt.Println(vals1)
	fmt.Println("-----------")
	// 二维数组返回
	vals2 := levelOrder(structures.Ints2TreeNode(arr))
	fmt.Println(vals2)
}
