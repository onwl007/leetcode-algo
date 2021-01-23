package main

import (
	"fmt"
	"testing"

	"github.com/onwl007/leetcode-algo/structures"
	"github.com/stretchr/testify/assert"
)

type Case struct {
	in  *ListNode
	out bool
}

func TestIsPalindrome(t *testing.T) {
	qs := []Case{
		{in: structures.Ints2ListNode([]int{1, 2}), out: false},
		{in: structures.Ints2ListNode([]int{1, 2, 2, 1}), out: true},
	}

	ast := assert.New(t)

	for i, v := range qs {
		ast.Equal(v.out, isPalindrome(v.in), fmt.Sprintf("第 %d 个测试用例报错", i))
		ast.Equal(v.out, isPalindrome1(v.in), fmt.Sprintf("第 %d 个测试用例报错", i))
	}
}
