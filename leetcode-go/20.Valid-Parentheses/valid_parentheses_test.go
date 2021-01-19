package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	in  string
	out bool
}

func Test_IsValid(t *testing.T) {
	qs := []Case{
		{in: "()", out: true},
		{in: "()[]{}", out: true},
		{in: "(]", out: false},
		{in: "{[]}", out: true},
	}
	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, isValid(v.in), "测试运行结果是否等于期望")
		ast.Equal(v.out, isValid1(v.in), "测试运行结果是否等于期望")
	}
}
