package main

import "sort"

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	s := []int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]]++
		} else {
			m[nums[i]] = 1
			// 记录不同的元素，去除重复元素
			s = append(s, nums[i])
		}
	}

	// 按照 map 里记录的次数从大到小排序
	sort.Slice(s, func(i, j int) bool {
		return m[s[i]] > m[s[j]]
	})

	return s[:k]
}
