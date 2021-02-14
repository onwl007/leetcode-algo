package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums []int
	k    int
	out  int
}

func TestFindKthLargest(t *testing.T) {
	qs := []Case{
		{nums: []int{3, 2, 1, 5, 6, 4}, k: 2, out: 5},
		{nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4, out: 4},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, findKthLargest(v.nums, v.k), "数组中的第 K 个最大元素")
	}
}
