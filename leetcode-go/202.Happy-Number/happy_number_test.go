package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	n   int
	out bool
}

func TestIsHappy(t *testing.T) {
	qs := []Case{
		{n: 19, out: true},
		{n: 2, out: false},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, isHappy(v.n), "快乐数")
		ast.Equal(v.out, isHappy1(v.n), "快乐数")
	}
}
