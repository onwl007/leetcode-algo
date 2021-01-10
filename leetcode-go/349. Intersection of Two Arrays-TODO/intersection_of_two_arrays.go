package main

// 遍历两个数组，存储的集合
// 然后循环长度短的集合，找到元素是否在第二个集合中
func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)

	for _, v := range nums1 {
		m1[v] = v
	}

	for _, v := range nums2 {
		m2[v] = v
	}

	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}

	var vals []int
	for _, v := range m1 {
		if _, ok := m2[v]; ok {
			vals = append(vals, v)
		}
	}
	return vals
}
