package structures

// 冒泡排序
func bubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		// 提前退出冒泡循环的标志位
		flag := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
				flag = true // 表示有数据交换
			}
		}
		if !flag {
			break // 没有数据交换，提前退出
		}
	}
}

// 插入排序
func insertSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		j := i - 1
		// 查找插入位置
		for ; j >= 0; j-- {
			if nums[j] > val {
				// 数据移动
				nums[j+1] = val
			} else {
				break
			}
		}
		// 插入数据
		nums[j+1] = val
	}
}
