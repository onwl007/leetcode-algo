package main

import "fmt"

func main() {
	lengthOfLongestSubstring("abcabcbb")
}

func lengthOfLongestSubstring(s string) int {
	fmt.Println(s)
	length := len(s)
	i, j := 0, 0
	maxlen := 0 // 存储当前窗口最长子串的长度
	for i = 0; i < length; i++ {
		for j = i + 1; j < length; j++ {
			fmt.Println(string(s[i]))
			if s[i] == s[j] {
				i++
				break
			} else {
				maxlen = j - i + 1
			}
		}
		maxlen = max(maxlen, j-i+1)
	}
	fmt.Println("max", maxlen)
	return maxlen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
