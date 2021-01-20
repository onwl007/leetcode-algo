package main

import "strings"

/*
 * @lc app=leetcode.cn id=1021 lang=golang
 *
 * [1021] 删除最外层的括号
 */

// @lc code=start
func removeOuterParentheses(S string) string {
	stack := []byte{}
	var res strings.Builder
	start := 1
	for i := 0; i < len(S); i++ {
		if S[i] == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				res.WriteString(S[start:i])
				start = i + 2
			}
		} else {
			stack = append(stack, S[i])
		}
	}
	return res.String()
}

// @lc code=end
