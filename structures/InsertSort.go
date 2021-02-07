package structures

func InsertSort(nums []int) {
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
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		// 插入数据
		nums[j+1] = val
	}
}
