package main

// 暴力解法
// 每次只跟元素的后面的元素比较大小
// 当前元素比后面一个元素的要小时，直接记录索引的差值即可，然后立刻退出第二重循环
// 进行下一个元素的比较
func dailyTemperatures(T []int) []int {
	vals := make([]int, len(T))

	for i := 0; i < len(T); i++ {
		if T[i] < 100 {
			for j := i + 1; j < len(T); j++ {
				if T[i] < T[j] {
					vals[i] = j - i
					break
				}
			}
		}
	}
	return vals
}
