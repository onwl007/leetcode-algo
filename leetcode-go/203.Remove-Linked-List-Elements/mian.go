package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

func main() {
	arr := []int{1, 2, 6, 3, 4, 5, 6}
	res := removeElements(structures.Ints2ListNode(arr), 6)
	fmt.Printf("-------- res %v", structures.ListNode2Ints(res))
}
