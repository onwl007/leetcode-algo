package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	n   int
	out bool
}

func TestIsPowerOfTwo(t *testing.T) {
	qs := []Case{
		{n: 1, out: true},
		{n: 16, out: true},
		{n: 218, out: false},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, isPowerOfTwo(v.n), "2 的幂")
	}
}
