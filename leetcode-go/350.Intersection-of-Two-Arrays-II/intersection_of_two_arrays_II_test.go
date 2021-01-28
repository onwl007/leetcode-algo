package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums1 []int
	nums2 []int
	out   []int
}

func TestIntersect(t *testing.T) {
	qs := []Case{
		{nums1: []int{1, 2, 2, 1}, nums2: []int{2, 2}, out: []int{2, 2}},
		{nums1: []int{4, 9, 5}, nums2: []int{9, 4, 9, 8, 4}, out: []int{4, 9}},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.ElementsMatch(v.out, intersect(v.nums1, v.nums2), "两个数组的交集 II")
		ast.ElementsMatch(v.out, intersect1(v.nums1, v.nums2), "两个数组的交集 II")
	}
}
