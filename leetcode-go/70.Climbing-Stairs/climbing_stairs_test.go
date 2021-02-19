package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	n   int
	out int
}

func TestClimbingStairs(t *testing.T) {
	qs := []Case{
		{n: 2, out: 2},
		{n: 3, out: 3},
		{n: 10, out: 89},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, climbStairs(v.n), "爬楼梯")
		ast.Equal(v.out, climbStairs1(v.n), "爬楼梯")
	}
}
