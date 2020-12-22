package main

import (
	"fmt"

	"github.com/onwl007/leetcode-algo/structures"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	node := connect(structures.Ints2Node(arr))
	fmt.Println(node)
}
