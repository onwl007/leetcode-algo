package main

func main() {
	lengthOfLongestSubstring("abcabcbb")
}

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	maxlen := 0
	l, r := 0, 0
	for ; r < length; r++ {
		for i := l; i < r; i++ {
			if s[r] == s[i] {
				l = i + 1
				break
			}
		}
		if (r - l + 1) > maxlen {
			maxlen = r - l + 1
		}
	}
	return maxlen
}

func lengthOfLongestSubstring2(s string) int {
	maxlen := 0
	length := len(s)
	m := make(map[byte]int)
	l := 0
	for r := 0; r < length; r++ {
		if _, ok := m[s[r]]; ok {
			l = max(m[s[r]], l)
		}
		maxlen = max(maxlen, r-l+1)
		m[s[r]] = r + 1
	}
	return maxlen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
