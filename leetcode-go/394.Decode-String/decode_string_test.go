package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	s   string
	out string
}

func TestDecodeString(t *testing.T) {
	qs := []Case{
		{s: "3[a]2[bc]", out: "aaabcbc"},
		{s: "3[a2[c]]", out: "accaccacc"},
		{s: "2[abc]3[cd]ef", out: "abcabccdcdcdef"},
		{s: "abc3[cd]xyz", out: "abccdcdcdxyz"},
	}

	ast := assert.New(t)
	for _, v := range qs {
		ast.Equal(v.out, decodeString(v.s), "字符串解码")
	}
}
