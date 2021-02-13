package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums []int
	out  int
}

func TestMajorityElement(t *testing.T) {
	qs := []Case{
		{nums: []int{3, 2, 3}, out: 3},
		{nums: []int{8, 8, 7, 7, 7}, out: 7},
		{nums: []int{2, 2, 1, 1, 1, 2, 2}, out: 2},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, majorityElement(v.nums), "多数元素")
		ast.Equal(v.out, majorityElement1(v.nums), "多数元素")
	}
}
