package main

import "sort"

// 排序
// 存在两种情况，都是正数时，乘积最大的是后三个的数的乘积
// 当存在负数时，两个最小的负数乘上最大的正数是最大的乘积·
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	a := nums[0] * nums[1] * nums[len(nums)-1]
	b := nums[len(nums)-3] * nums[len(nums)-2] * nums[len(nums)-1]
	return max(a, b)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
