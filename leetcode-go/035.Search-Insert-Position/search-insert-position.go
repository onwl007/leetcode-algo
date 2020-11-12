package main

func searchInsert(nums []int, target int) int {
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
