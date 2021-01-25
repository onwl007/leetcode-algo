package main

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	return []int{bsearchLeft(nums, target), bsearchRight(nums, target)}
}

// 二分查找找第一个等于的元素的位置
func bsearchLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else if target == nums[mid] {
			if mid == 0 || nums[mid-1] != target {
				return mid
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func bsearchRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else if target == nums[mid] {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

// @lc code=end

func leftBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// 搜索区间 [left, right]
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			// 搜索区间变为 [mid+1, right]
			left = mid + 1
		} else if nums[mid] > target {
			// 搜索区间变为 [left, mid-1]
			right = mid - 1
		} else if nums[mid] == target {
			// 搜索右侧边界
			right = mid - 1
		}
	}

	// 检查出界情况
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

func rightBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 收缩左侧边界
			left = mid + 1
		}
	}
	// 这里改为检查 right 越界的情况
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

// 这样也能实现，但不符合题目时间复杂度为 Ologn 的要求
func searchRange1(nums []int, target int) []int {
	start, end := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			start = i
			break
		}
	}

	for j := len(nums) - 1; j >= 0; j-- {
		if nums[j] == target {
			end = j
			break
		}
	}
	return []int{start, end}
}
