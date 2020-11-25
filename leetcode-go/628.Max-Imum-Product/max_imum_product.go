package main

import (
	"math"
	"sort"
)

func maximumProduct(nums []int) int {
	min1, min2 := math.MaxInt64, math.MaxInt64
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64

	for _, v := range nums {
		if v <= min1 {
			min2 = min1
			min1 = v
		} else if v <= min2 {
			min2 = v
		}
		if v >= max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v >= max2 {
			max3 = max2
			max2 = v
		} else if v >= max3 {
			max3 = v
		}
	}

	a := min1 * min2 * max1
	b := max1 * max2 * max3
	return max(a, b)
}

// 排序
// 存在两种情况，都是正数时，乘积最大的是后三个的数的乘积
// 当存在负数时，两个最小的负数乘上最大的正数是最大的乘积·
func maximumProduct1(nums []int) int {
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
