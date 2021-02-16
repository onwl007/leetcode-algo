package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums []int
	out  bool
}

func TestContainsDuplicate(t *testing.T) {
	qs := []Case{
		{nums: []int{1, 2, 3, 1}, out: true},
		{nums: []int{1, 2, 3, 4}, out: false},
		{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, out: true},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, containsDuplicate(v.nums), "存在重复元素")
		ast.Equal(v.out, containsDuplicate1(v.nums), "存在重复元素")
	}
}
