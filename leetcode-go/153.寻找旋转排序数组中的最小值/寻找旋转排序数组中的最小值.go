package main

import "fmt"

func main() {
	n := findMin([]int{4, 5, 6, 7, 0, 1, 2})
	fmt.Println("res: ", n)
}

/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

// @lc code=end
