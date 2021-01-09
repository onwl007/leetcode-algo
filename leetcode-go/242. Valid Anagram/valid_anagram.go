package main

import "sort"

// 哈希表解法
// 先循环 s 将每个字母的次数记录下来
// 然后循环 t，从哈希表中取值，如果有则减 1，没有的字母次数置为 1
// 最后判断哈希表的长度
func isAnagram(s string, t string) bool {
	ms := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if _, ok := ms[s[i]]; ok {
			ms[s[i]]++
		} else {
			ms[s[i]] = 1
		}
	}

	for i := 0; i < len(t); i++ {
		if _, ok := ms[t[i]]; ok {
			ms[t[i]]--
		} else {
			ms[t[i]] = 1
		}

		if ms[t[i]] == 0 {
			delete(ms, t[i])
		}
	}

	if len(ms) != 0 {
		return false
	}
	return true
}

// 排序解法
// 排序后字符串相等则是异位词
func isAnagram1(s string, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})
	return string(s1) == string(s2)
}
