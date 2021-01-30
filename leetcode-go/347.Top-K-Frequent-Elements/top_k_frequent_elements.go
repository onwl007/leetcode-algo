package main

import (
	"fmt"
	"sort"
)

func main() {
	vals := topKFrequent([]int{4, 4, 4, 2, 2, 3, 1}, 2)
	fmt.Println(vals)
}

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	vals := []int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]]++
		} else {
			m[nums[i]] = 1
			vals = append(vals, nums[i])
		}
	}

	sort.Slice(vals, func(i, j int) bool {
		return m[vals[i]] > m[vals[j]]
	})

	return vals[:k]
}

// @lc code=end
