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
