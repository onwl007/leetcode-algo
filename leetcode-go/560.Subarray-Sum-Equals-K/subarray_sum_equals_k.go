package main

import "fmt"

func main() {
	res := subarraySum([]int{1, 1, 1}, 2)
	fmt.Println(res)
}

// TODO: 前缀和解法
/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}

// @lc code=end
