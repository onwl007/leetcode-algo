package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Pas struct {
	nums   []int
	target int
}

type Case struct {
	in  Pas
	out []int
}

func Test_Problem34(t *testing.T) {
	qs := []Case{
		{in: Pas{nums: []int{5, 7, 7, 8, 8, 10}, target: 8}, out: []int{3, 4}},
		{in: Pas{nums: []int{5, 7, 7, 8, 8, 10}, target: 6}, out: []int{-1, -1}},
	}

	ast := assert.New(t)

	for _, v := range qs {
		ast.Equal(v.out, searchRange(v.in.nums, v.in.target), "在排序数组中查找元素的第一个和最后一个位置")
		ast.Equal(v.out, searchRange1(v.in.nums, v.in.target), "在排序数组中查找元素的第一个和最后一个位置")
	}
}
