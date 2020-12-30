package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

func main() {
	arr := []int{1, 1, 2, 1}
	res := isPalindrome(structures.Ints2ListNode(arr))
	dfs(structures.Ints2ListNode(arr))
	fmt.Println(res)
}
