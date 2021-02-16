package main

import "sort"

// 哈希表
func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, v := range nums {
		m[v] = struct{}{}
	}

	return len(nums) != len(m)
}

// 排序
// 排完序后，相等的元素一定相邻
// 注意数组越界
func containsDuplicate1(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
