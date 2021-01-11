package main

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}

	var vals []int
	for _, v := range nums2 {
		if m[v] > 0 {
			if _, ok := m[v]; ok {
				vals = append(vals, v)
				m[v]--
			}
		}
	}
	return vals
}

// 双指针
// 对数组排序后使用双指针
func intersect1(nums1 []int, nums2 []int) []int {
	var vals []int

	sort.Ints(nums1)
	sort.Ints(nums2)
	i1, i2 := 0, 0
	l1, l2 := len(nums1), len(nums2)

	for i1 < l1 && i2 < l2 {
		if nums1[i1] < nums2[i2] {
			i1++
		} else if nums1[i1] < nums2[i2] {
			i2++
		} else {
			vals = append(vals, nums1[i1])
			i1++
			i2++
		}
	}
	return vals
}
