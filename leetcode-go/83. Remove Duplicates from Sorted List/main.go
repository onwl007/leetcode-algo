package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

func main() {
	arr := []int{1, 1, 2}
	res := deleteDuplicates(structures.Ints2ListNode(arr))
	fmt.Printf("------- res %v\n", structures.ListNode2Ints(res))
}
