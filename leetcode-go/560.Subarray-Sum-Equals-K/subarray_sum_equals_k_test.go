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

func TestSubarraySum(t *testing.T) {
	qs := []Case{
		{nums: []int{1, 1, 1}, k: 2, out: 2},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, subarraySum(v.nums, v.k), "和为K的子数组")
	}
}
