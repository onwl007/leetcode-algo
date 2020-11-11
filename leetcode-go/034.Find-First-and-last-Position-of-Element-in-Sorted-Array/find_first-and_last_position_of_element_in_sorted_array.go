package main

func searchRange(nums []int, target int) []int {
	return []int{searchFistEqual(nums, target), searchLastEqual(nums, target)}
}

// 二分查找找第一个等于的元素的位置
func searchFistEqual(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-hi)/2
		if target < nums[mid] {
			hi = mid - 1
		} else if target > nums[mid] {
			lo = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				return mid
			}
			hi = mid - 1
		}
	}
	return -1
}

// 二分查找最后一个等于的元素的位置
func searchLastEqual(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-hi)/2
		if target < nums[mid] {
			hi = mid - 1
		} else if target > nums[mid] {
			lo = mid + 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				return mid
			}
			lo = mid + 1
		}
	}
	return -1
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
