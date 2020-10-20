package main

func main() {
	lengthOfLongestSubstring("abcabcbb")
}

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	maxlen := 0
	left, right, tmp := 0, 0, 0
	for left = 0; left < length; left++ {
		for right = tmp; right < left; right++ {
			if s[right] == s[left] {
				tmp = right + 1
				break
			}
		}
		if (left - tmp + 1) > maxlen {
			maxlen = left - tmp + 1
		}
	}
	return maxlen
}
