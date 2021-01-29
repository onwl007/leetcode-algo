package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums   []int
	target int
	out    int
}

func Test_Problem35(t *testing.T) {
	qs := []Case{
		{nums: []int{1, 3, 5, 6}, target: 5, out: 2},
		{nums: []int{1, 3, 5, 6}, target: 2, out: 1},
		{nums: []int{1, 3, 5, 6}, target: 7, out: 4},
		{nums: []int{1, 3, 5, 6}, target: 0, out: 0},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, searchInsert(v.nums, v.target), "搜索插入位置")
		ast.Equal(v.out, searchInsert1(v.nums, v.target), "搜索插入位置")
	}

}
