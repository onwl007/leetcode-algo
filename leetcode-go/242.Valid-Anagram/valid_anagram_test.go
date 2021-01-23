package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	in1 string
	in2 string
	out bool
}

func TestIsAnagram(t *testing.T) {
	qs := []Case{
		{in1: "anagram", in2: "nagaram", out: true},
		{in1: "rat", in2: "car", out: false},
		{in1: "a", in2: "ab", out: false},
		{in1: "aaaaaaaa", in2: "aaa", out: false},
	}

	ast := assert.New(t)

	for i, v := range qs {
		ast.Equal(v.out, isAnagram(v.in1, v.in2), fmt.Sprintf("第 %d 个测试用例出错， 有效的字母异位词", i+1))
		ast.Equal(v.out, isAnagram1(v.in1, v.in2), fmt.Sprintf("第 %d 个测试用例出错， 有效的字母异位词", i+1))
	}
}
