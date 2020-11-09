package main

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			j++
		}
	}
	return i
}

// 当数组包含很少要删除的元素的情况，方法二可以减少对多余的数组元素复制操作
func removeElement1(nums []int, val int) int {
	i := 0
	n := len(nums)
	for i < n {
		// 当前元素与 val 相等时，将最后一个元素的值复制到当前位置
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}
	return n
}
