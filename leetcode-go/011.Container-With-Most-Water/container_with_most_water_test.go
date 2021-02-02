package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	height []int
	out    int
}

func Test_Problem11(t *testing.T) {
	qs := []Case{
		{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, out: 49},
		{height: []int{1, 1}, out: 1},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, maxArea(v.height), "盛最多水的容器")
		ast.Equal(v.out, maxArea1(v.height), "盛最多水的容器")
	}
}
