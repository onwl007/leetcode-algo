package main

// 二分查找
func search(nums []int, target int) int {
	mid, left, right := 0, 0, len(nums)-1
	// 考虑只有一个元素的情况
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		// 左侧是排好序且递增的
		if nums[left] <= nums[mid] {
			// 在 [left, mid)区间内查找
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 在 [mid, right]区间内查找
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// 时间复杂度必须是 logN，所以这种方法不符和题意
func search1(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}
