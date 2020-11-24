package main

import "sort"

// 时间复杂度太高
// 去重
// 排序
// 步骤太多
func thirdMax1(nums []int) int {
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
