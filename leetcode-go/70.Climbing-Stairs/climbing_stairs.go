package main

/*
 * 递归为啥超时呢，因为有大量的重复计算，可以采用增加缓存的方式，涉及重复计算的地方，
 * 直接从缓存中取
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start

var m map[int]int

func climbStairs(n int) int {
	m = make(map[int]int)
	return fn(n)
}

func fn(n int) int {
	if n <= 2 {
		return n
	} else if v, ok := m[n]; ok {
		return v
	}

	res := fn(n-1) + fn(n-2)
	m[n] = res
	return res
}

// @lc code=end

// 递归超时
func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}
