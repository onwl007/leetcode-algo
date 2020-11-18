package main

func findDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		count := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				count++
			}
		}

		if count > mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
