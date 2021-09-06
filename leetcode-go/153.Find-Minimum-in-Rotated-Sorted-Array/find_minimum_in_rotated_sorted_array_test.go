package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	in  []int
	out int
}

func TestFindMin(t *testing.T) {
	qs := []Case{
		{in: []int{3, 4, 5, 1, 2}, out: 1},
		{in: []int{4, 5, 6, 7, 0, 1, 2}, out: 0},
		{in: []int{11, 13, 15, 17}, out: 11},
		{in: []int{1}, out: 1},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, findMin(v.in), "")
	}
}
