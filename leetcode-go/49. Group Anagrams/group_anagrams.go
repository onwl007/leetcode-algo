package main

import "sort"

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	var vals [][]string
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		m[string(s)] = append(m[string(s)], str)
	}

	for _, v := range m {
		vals = append(vals, v)
	}
	return vals
}

// @lc code=end
