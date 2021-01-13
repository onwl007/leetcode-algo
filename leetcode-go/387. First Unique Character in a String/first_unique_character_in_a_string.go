package main

func firstUniqChar(s string) int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; ok && m[s[i]] == 1 {
			return i
		}
	}
	return -1
}
