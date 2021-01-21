package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	in  []int
	out int
}

func TestMaximumProduct(t *testing.T) {
	qs := []Case{
		{in: []int{1, 2, 3}, out: 6},
		{in: []int{1, 2, 3, 4}, out: 24},
		{in: []int{-1, -2, -3}, out: -6},
	}
	ast := assert.New(t)

	for i, v := range qs {
		ast.Equal(v.out, maximumProduct(v.in), fmt.Sprintf("第 %d 个测试用例，三个数最大乘积出错", i))
		ast.Equal(v.out, maximumProduct1(v.in), fmt.Sprintf("第 %d 个测试用例，三个数最大乘积出错", i))
	}
}
