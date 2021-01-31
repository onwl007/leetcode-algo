package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums []int
	out  int
}

func TestThirdMax(t *testing.T) {
	qs := []Case{
		{nums: []int{3, 2, 1}, out: 1},
		{nums: []int{1, 2}, out: 2},
		{nums: []int{2, 2, 3, 1}, out: 1},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, thirdMax(v.nums), "第三大的数")
		ast.Equal(v.out, thirdMax1(v.nums), "第三大的数")
		ast.Equal(v.out, thirdMax2(v.nums), "第三大的数")
	}
}
