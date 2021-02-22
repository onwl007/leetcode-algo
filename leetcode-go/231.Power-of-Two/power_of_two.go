package main

/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	for n%2 == 0 {
		n /= 2
	}
	return n == 1
}

// @lc code=end
