package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	res := reverseList(structures.Ints2ListNode(arr))
	fmt.Printf("------- reverse %v\n", structures.ListNode2Ints(res))
}
