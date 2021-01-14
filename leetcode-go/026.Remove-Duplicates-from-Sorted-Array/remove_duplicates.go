package main

import "fmt"

func main() {
	arr1 := []int{1, 1, 2}
	arr2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(arr1))
	fmt.Println(removeDuplicates(arr2))
}

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
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

// @lc code=end
