package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums []int
	out  []int
}

func TestFindErrorNums(t *testing.T) {
	qs := []Case{
		{nums: []int{1, 2, 2, 4}, out: []int{2, 3}},
		{nums: []int{1, 1}, out: []int{1, 2}},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.ElementsMatch(v.out, findErrorNums(v.nums), "findErrorNums 错误的集合")
		ast.ElementsMatch(v.out, findErrorNums1(v.nums), "findErrorNums1 错误的集合")
		ast.ElementsMatch(v.out, findErrorNums2(v.nums), "findErrorNums2 错误的集合")
		ast.ElementsMatch(v.out, findErrorNums3(v.nums), "findErrorNums3 错误的集合")
	}
}
