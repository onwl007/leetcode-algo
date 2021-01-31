package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums   []int
	target int
	out    []int
}

func TestTwoSum(t *testing.T) {
	qs := []Case{
		{nums: []int{3, 2, 4}, target: 6, out: []int{1, 2}},
		{nums: []int{3, 2, 4}, target: 5, out: []int{0, 1}},
		{nums: []int{0, 8, 7, 3, 3, 4, 2}, target: 11, out: []int{1, 3}},
		{nums: []int{0, 1}, target: 1, out: []int{0, 1}},
		{nums: []int{0, 3}, target: 5, out: []int{}},
		{nums: []int{3, 3}, target: 6, out: []int{0, 1}},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.ElementsMatch(v.out, twoSum(v.nums, v.target), "两数之和")
		ast.ElementsMatch(v.out, twoSum1(v.nums, v.target), "两数之和")
	}
}
