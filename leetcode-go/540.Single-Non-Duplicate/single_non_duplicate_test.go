package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	in  []int
	out int
}

func TestSingleNonDuplicate(t *testing.T) {
	qs := []Case{
		{in: []int{1, 1, 2, 3, 3, 4, 4, 8, 8}, out: 2},
		{in: []int{3, 3, 7, 7, 10, 11, 11}, out: 10},
	}

	ast := assert.New(t)

	for _, v := range qs {
		ast.Equal(v.out, singleNonDuplicate(v.in), "有序数组中的单一元素")
		ast.Equal(v.out, singleNonDuplicate1(v.in), "有序数组中的单一元素")
	}
}
