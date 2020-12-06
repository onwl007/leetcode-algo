package main

// 暴力法
// 遍历 1 到 n 的所有数字，检查每个数在数组出现的次数，或未出现
func findErrorNums(nums []int) []int {
	dup, missing := -1, -1
	for i := 1; i <= len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] == i {
				count++
			}
		}
		if count == 2 {
			dup = i
		} else if count == 0 {
			missing = i
		}
	}
	return []int{dup, missing}
}
