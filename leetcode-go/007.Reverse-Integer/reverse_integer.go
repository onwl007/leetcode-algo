package main

import "math"

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	if x == math.MinInt32 {
		return 0
	}

	if x < 0 {
		return -reverse(-x)
	}

	n := 0
	for x != 0 {
		if n > (math.MaxInt32-x%10)/10 {
			return 0
		}
		n = n*10 + x%10
		x /= 10
	}
	return n
}

// @lc code=end
