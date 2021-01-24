package main

/*
 * @lc app=leetcode.cn id=540 lang=golang
 *
 * [540] 有序数组中的单一元素
 */

// @lc code=start
func singleNonDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		mod := (right-mid)%2 == 0
		if nums[mid+1] == nums[mid] {
			if mod {
				left = mid + 2
			} else {
				right = mid - 1
			}
		} else if nums[mid-1] == nums[mid] {
			if mod {
				right = mid - 2
			} else {
				left = mid + 1
			}
		} else {
			return nums[mid]
		}
	}
	return nums[left]
}

// @lc code=end

// 线性搜索，不符合题意
func singleNonDuplicate1(nums []int) int {
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}
