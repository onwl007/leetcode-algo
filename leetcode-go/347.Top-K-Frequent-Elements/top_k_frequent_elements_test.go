package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums []int
	k    int
	out  []int
}

func TestTopKFrequent(t *testing.T) {
	qs := []Case{
		{nums: []int{1, 1, 1, 2, 2, 3}, k: 2, out: []int{1, 2}},
		{nums: []int{1}, k: 1, out: []int{1}},
		{nums: []int{4, 4, 4, 2, 2, 3, 1}, k: 2, out: []int{4, 2}},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.ElementsMatch(v.out, topKFrequent(v.nums, v.k), "前 K 个高频元素")
	}
}
