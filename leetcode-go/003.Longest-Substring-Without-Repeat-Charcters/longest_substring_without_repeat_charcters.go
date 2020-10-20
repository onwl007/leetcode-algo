package leetcode

func lengthOfLongestSubstring(s string) int {
	s = "abcabcbb"
	m := make(map[byte]int)
	length := len(s)
	maxlen := 0 // 存储当前窗口最长子串的长度
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if _, ok := m[s[j]]; ok {
				break
			} else {
				m[s[j]] = j
			}
		}
		maxlen = max(maxlen, len(m))
		if maxlen >= length-1-i {
			break
		}
	}
	return maxlen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
