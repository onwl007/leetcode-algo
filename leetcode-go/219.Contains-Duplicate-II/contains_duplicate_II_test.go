package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums []int
	k    int
	out  bool
}

func TestContainsNearbyDuplicate(t *testing.T) {
	qs := []Case{
		{nums: []int{1, 2, 3, 1}, k: 3, out: true},
		{nums: []int{1, 0, 1, 1}, k: 1, out: true},
		{nums: []int{1, 2, 3, 1, 2, 3}, k: 2, out: false},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, containsNearbyDuplicate(v.nums, v.k), "存在重复元素 II")
		ast.Equal(v.out, containsNearbyDuplicate1(v.nums, v.k), "存在重复元素 II")
		ast.Equal(v.out, containsNearbyDuplicate2(v.nums, v.k), "存在重复元素 II")
	}
}
