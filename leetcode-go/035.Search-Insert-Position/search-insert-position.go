package main

// 这道题是典型的二分查找的变形应用
// 查找最后一个小于等于的元素
/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if target < nums[mid] {
			right = mid - 1
		} else if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			if mid == len(nums)-1 || nums[mid+1] > target {
				return mid + 1
			} else {
				left = mid + 1
			}
		}
	}
	return 0
}

// @lc code=end

func searchInsert1(nums []int, target int) int {
	x := 0
	for i := 0; i < len(nums); i++ {
		if target > nums[i] {
			x++
		} else {
			return i
		}
	}
	return x
}
