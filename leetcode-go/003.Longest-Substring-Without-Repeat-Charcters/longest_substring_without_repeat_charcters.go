package main

import "fmt"

func main() {
	lengthOfLongestSubstring("abcabcbb")
}

func lengthOfLongestSubstring(s string) int {
	fmt.Println(s)
	length := len(s)
	i, j, k := 0, 0, 0
	maxlen := 0 // 存储当前窗口最长子串的长度
	for j = 0; j < length; j++ {
		for k = i; k < j; k++ {
			fmt.Println(string(s[j]))
			if s[k] == s[j] {
				i = k + 1
				break
			}
		}
		if (j - i + 1) > maxlen {
			maxlen = j - i + 1
		}
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
