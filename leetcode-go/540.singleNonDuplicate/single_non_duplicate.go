package main

// 有序数组中的单一元素
// 线性搜索暴力解法
func singleNonDuplicate1(nums []int) int {
	for i := 0; i < len(nums); i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}
