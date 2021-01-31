package main

import (
	"math"
	"sort"
)

// 动态维护 3 个最大值即可，注意数组有重复数据的情况。
// 如果只有 2 个数或者 1 个数，则返回最大值即可。
/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 */

// @lc code=start
func thirdMax(nums []int) int {
	one, two, three := math.MinInt64, math.MinInt64, math.MinInt64
	for _, v := range nums {
		if v > one {
			three = two
			two = one
			one = v
		} else if v < one && v > two {
			three = two
			two = v
		} else if v < two && v > three {
			three = v
		}
	}
	if three == math.MinInt64 {
		return one
	}
	return three
}

// @lc code=end

func thirdMax1(nums []int) int {
	m := make(map[int]struct{})
	s := make([]int, 0, len(nums))
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v] = struct{}{}
		} else {
			m[v] = struct{}{}
			s = append(s, v)
		}
	}
	sort.Ints(s)
	if len(s) <= 2 {
		return s[len(s)-1]
	}
	return s[len(s)-3]
}

// 时间复杂度太高
// 去重
// 排序，时间复杂度不是 O(n)，所以不符合题意
// 步骤太多
func thirdMax2(nums []int) int {
	nums = removeDuplicate(nums)
	sort.Ints(nums)
	if len(nums) <= 2 {
		return nums[len(nums)-1]
	}
	return nums[len(nums)-3]
}

func removeDuplicate(nums []int) []int {
	result := make([]int, 0, len(nums))
	temp := map[int]struct{}{}
	for _, v := range nums {
		if _, ok := temp[v]; !ok {
			temp[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}
