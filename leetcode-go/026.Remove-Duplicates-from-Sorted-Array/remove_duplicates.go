package main

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 通过双指针比较法，将数组前面的元素都赋值成不相同的，只需要返回 i 指针的下一位即可
	i := 0
	for j := 1; j < n; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
