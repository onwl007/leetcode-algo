package structures

// 冒泡排序
func BubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
			}
		}
	}
}

// 改进版冒泡排序
func BubbleSortImprove(nums []int) {
	for i := 0; i < len(nums); i++ {
		// 提前退出冒泡循环的标志位
		flag := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
				flag = true // 表示有数据交换
			}
		}
		if !flag {
			break // 没有数据交换，提前退出
		}
	}
}
