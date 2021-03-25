package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	s   string
	p   string
	out []int
}

func TestFindAnagrams(t *testing.T) {
	qs := []Case{
		{s: "cbaebabacd", p: "abc", out: []int{0, 6}},
		{s: "abab", p: "ab", out: []int{0, 1, 2}},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, findAnagrams(v.s, v.p), "找到字符串中所有字母异位词")
	}
}
