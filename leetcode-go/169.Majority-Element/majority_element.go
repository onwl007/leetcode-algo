package main

import "sort"

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v > len(nums)/2 {
			return k
		}
	}
	return -1
}

// @lc code=end

func majorityElement1(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
