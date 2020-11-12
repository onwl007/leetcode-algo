package main

func searchInsert(nums []int, target int) int {
	x := 0
	for i := 0; i < len(nums); i++ {
		if target == nums[i] {
			return i
		} else if target < nums[i] {
			return i
		} else if target > nums[i] {
			x++
		}
	}
	return x
}
