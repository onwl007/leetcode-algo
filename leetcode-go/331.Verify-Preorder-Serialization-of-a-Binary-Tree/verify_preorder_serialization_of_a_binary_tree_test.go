package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	in  string
	out bool
}

func TestIsValidSerialization(t *testing.T) {
	ast := assert.New(t)
	qs := []Case{
		{in: "9,3,4,#,#,1,#,#,2,#,6,#,#", out: true},
		{in: "1,#", out: false},
		{in: "9,#,#,1", out: false},
		{in: "#,7,6,9,#,#,#", out: false},
		{in: "#,#,3,5,#", out: false},
	}

	for i, v := range qs {
		ast.Equal(v.out, isValidSerialization(v.in), fmt.Sprintf("isValidSerialization 第 %v 个测试用例执行错误", i+1))
		ast.Equal(v.out, isValidSerialization1(v.in), fmt.Sprintf("isValidSerialization1 第 %v 个测试用例执行错误", i+1))
	}
}
