package main

// 标准二分查找，适合有序数组
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		// mid := (left + right)/2
		mid := left + (right-left)/2 // 这样写是为了防止 left + right 整型溢出
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		}
	}
	return -1
}
