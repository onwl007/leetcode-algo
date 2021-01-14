package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}

/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	var vals []int
	var win [26]int
	var hash [26]int

	for i := 0; i < len(p); i++ {
		hash[p[i]-'a']++
	}

	l, r := 0, 0
	for ; r < len(s); r++ {
		win[s[r]-'a']++
		if r+1 < len(p) {
			continue
		}
		if win == hash {
			vals = append(vals, l)
		}
		win[s[l]-'a']--
		l++
	}
	return vals
}

// @lc code=end

// 排序法
// 先对 p 排序
// 然后再 s 中每次循环选取 p 长度的字符排序，每次循环一次往前进 1
// 超时解法
func findAnagrams1(s string, p string) []int {
	n := len(p)
	sp := []byte(p)
	sort.Slice(sp, func(i, j int) bool {
		return sp[i] < sp[j]
	})
	var vals []int
	sb := []byte(s)
	for i := 0; i < len(sb)-n; i++ {
		b := sb[i : i+n]
		sc := make([]byte, len(b))
		copy(sc, b)
		sort.Slice(sc, func(i, j int) bool {
			return sc[i] < sc[j]
		})
		if string(sc) == string(sp) {
			vals = append(vals, i)
		}
	}
	return vals
}
