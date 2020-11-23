package main

func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			maxCount = max(maxCount, count)
			count = 0
		}
	}
	return max(maxCount, count)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
