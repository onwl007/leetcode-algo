package leetcode

func lengthOfLongestSubstring(s string) int {
	s = "abcabcbb"
	max := 0
	for i := 0; i < len(s); i++ {
		for j := i; i < len(s); j++ {
			if s[i] == s[j] {
				break
			}
			if j-i+1 > max {
				max = j - i + 1
			}
		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
