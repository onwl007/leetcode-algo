package main

// 使用 map
func findErrorNums(nums []int) []int {
	res := []int{}
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[v]; ok {
			res = append(res, v)
		}
		m[v] = struct{}{}
	}
	for i, _ := range nums {
		if _, ok := m[i+1]; !ok {
			res = append(res, i+1)
		}
	}
	return res
}

// 暴力法
// 遍历 1 到 n 的所有数字，检查每个数在数组出现的次数，或未出现
func findErrorNums1(nums []int) []int {
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

// 优化暴力解法
// 当 dup 和 missing 两个变量都大于 0 时，可以直接退出循环
func findErrorNums2(nums []int) []int {
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
		if dup > 0 && missing > 0 {
			break
		}
	}
	return []int{dup, missing}
}
