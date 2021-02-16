package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
	out   []int
}

func TestMerge(t *testing.T) {
	qs := []Case{
		{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3, out: []int{1, 2, 2, 3, 5, 6}},
		{nums1: []int{1}, m: 1, nums2: []int{}, n: 0, out: []int{1}},
	}

	ast := assert.New(t)
	for _, v := range qs {
		merge(v.nums1, v.m, v.nums2, v.n)
		ast.Equal(v.out, v.nums1, "合并两个有序数组")
	}
}
