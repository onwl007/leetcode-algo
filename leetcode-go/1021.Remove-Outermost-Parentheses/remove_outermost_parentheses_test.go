package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	in  string
	out string
}

func TestRemoveOuterParentheses(t *testing.T) {
	qs := []Case{
		{in: "(()())(())", out: "()()()"},
		{in: "(()())(())(()(()))", out: "()()()()(())"},
		{in: "()()", out: ""},
	}
	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, removeOuterParentheses(v.in), "删除最外层的括号")
	}
}
