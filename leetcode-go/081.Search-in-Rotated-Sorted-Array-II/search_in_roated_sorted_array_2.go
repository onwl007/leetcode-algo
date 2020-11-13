package main

import "sort"

func search(nums []int, target int) bool {
	return false
}

// 不符合题意的解法，先进行了一次升序排序，然后再二分查找
func search1(nums []int, target int) bool {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if target == nums[mid] {
			return true
		} else if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		}
	}
	return false
}

// 直接遍历，时间复杂度是 O(n)，也能做出来，但是不符合题意
func search2(nums []int, target int) bool {
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return true
		}
	}
	return false
}
