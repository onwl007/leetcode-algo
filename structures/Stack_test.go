package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	ast := assert.New(t)

	s := NewStack()
	ast.True(s.IsEmpty(), "检查初始化的栈是否为空")

	start, end := 0, 100
	for i := 0; i < end; i++ {
		s.Push(i)
		ast.Equal(i-start+1, s.Length(), "入栈后检查栈的长度")
	}

	for i := end - 1; i >= start; i-- {
		ast.Equal(i, s.Pop(), "从栈中 pop 元素")
	}

	ast.True(s.IsEmpty(), "检查所有元素出栈后的栈的长度")
}
