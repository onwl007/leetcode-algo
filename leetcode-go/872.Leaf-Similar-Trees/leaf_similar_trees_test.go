package main

import (
	"fmt"
	"testing"

	"github.com/onwl007/leetcode-algo/structures"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	in1 *TreeNode
	in2 *TreeNode
	out bool
}

func TestLeafSimilar(t *testing.T) {
	arr1 := []int{3, 5, 1, 6, 2, 9, 8, structures.NULL, structures.NULL, 7, 4}
	arr2 := []int{3, 5, 1, 6, 7, 4, 2, structures.NULL, structures.NULL, structures.NULL, structures.NULL, structures.NULL, structures.NULL, 9, 8}
	qs := []Case{
		{in1: structures.Ints2TreeNode(arr1), in2: structures.Ints2TreeNode(arr2), out: true},
		{in1: structures.Ints2TreeNode([]int{1}), in2: structures.Ints2TreeNode([]int{1}), out: true},
		{in1: structures.Ints2TreeNode([]int{1}), in2: structures.Ints2TreeNode([]int{2}), out: false},
		{in1: structures.Ints2TreeNode([]int{1, 2}), in2: structures.Ints2TreeNode([]int{2, 2}), out: true},
	}

	ast := assert.New(t)

	for i, v := range qs {
		ast.Equal(v.out, leafSimilar(v.in1, v.in2), fmt.Sprintf("第 %d 个测试用例，叶子相似的树错误", i))
	}
}
