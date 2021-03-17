package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	x   int
	out int
}

func TestReverse(t *testing.T) {
	qs := []Case{
		{x: 123, out: 321},
		{x: -123, out: -321},
		{x: 120, out: 21},
		{x: 0, out: 0},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, reverse(v.x), "整数反转")
	}
}
