package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	res := findKthLargest(arr, k)
	fmt.Printf("arr ---- %v\n", arr)
	fmt.Printf("res %d\n", res)
}

/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	res := quickFind(nums, k, 0, len(nums)-1)
	fmt.Println(nums)
	return res
}

func quickFind(nums []int, k, left, right int) int {
	if left >= right {
		return nums[left]
	}
	p := partition(nums, left, right)
	if k == p+1 {
		return nums[p]
	} else if k < p+1 {
		return quickFind(nums, k, left, p-1)
	} else {
		return quickFind(nums, k, p+1, right)
	}
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] > pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

// @lc code=end

func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
