package main

// 异或解法
func singleNumber(nums []int) int {
	single := 0
	for i := 0; i < len(nums); i++ {
		single ^= nums[i]
	}
	return single
}

// 两次循环，解法不是很好
func singleNumber1(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]]++
		} else {
			m[nums[i]] = 1
		}
	}

	for k, v := range m {
		if v != 1 {
			continue
		}
		return k
	}
	return 0
}

// 双重循环，暴力法解决
func singleNumber2(nums []int) int {
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] {
				count++
			}
		}
		if count == 1 {
			return nums[i]
		}
	}
	return 0
}
