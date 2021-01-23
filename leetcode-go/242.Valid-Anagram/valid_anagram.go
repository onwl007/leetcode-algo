package main

import "sort"

// 哈希表解法
// 先循环 s 将每个字母的次数记录下来
// 然后循环 t，从哈希表中取值，如果有则减 1，没有的字母次数置为 1
// 最后判断哈希表的长度
/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		if _, ok := m[t[i]]; ok {
			m[t[i]]--
		} else {
			m[t[i]] = 1
		}

		if m[t[i]] == 0 {
			delete(m, t[i])
		}
	}
	return len(m) == 0
}

// @lc code=end

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
