package main

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	m := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != m[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// @lc code=end

func isValid1(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if (s[i] == '(') || (s[i] == '[') || (s[i] == '{') {
			stack = append(stack, s[i])
		} else if ((s[i] == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			((s[i] == ']') && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			(s[i] == '}' && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
