package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	res := decodeString("3[a]2[bc]")
	fmt.Printf("result: %v\n", res)
}

/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
func decodeString(s string) string {
	res := ""
	strStack := []string{}
	numStack := []int{}
	var count int
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num, _ := strconv.Atoi(string(s[i]))
			count = count*10 + num
		} else if s[i] == '[' {
			strStack = append(strStack, res)
			res = ""
			numStack = append(numStack, count)
			count = 0
		} else if s[i] == ']' {
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			num := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			res = string(str) + strings.Repeat(res, num)
		} else {
			res += string(s[i])
		}
	}
	return res
}

// @lc code=end
