package main

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		}
	}
	return left
}

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
