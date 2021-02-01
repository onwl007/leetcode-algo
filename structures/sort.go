package structures

// 冒泡排序
func BubbleSort(nums []int) {
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
				nums[j+1] = val
			} else {
				break
			}
		}
		// 插入数据
		nums[j+1] = val
	}
}

// 归并排序
func MergeSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	merge(nums, left, mid, right)
}

// 合并 [l, r] 两部分数据，mid 左半部分的终点，mid + 1 右半部分的起点
func merge(nums []int, left, mid, right int) {
	temp := make([]int, right-left+1) // 开辟一个临时数组，存放合并之后的数组
	i := left                         // 左边的指针，指向左边数组的第一个元素
	j := mid + 1                      // 右边的指针，指向右边数组的第一个元素
	k := 0                            // 用于记录临时数组当前指针指向
	// 一直遍历左右两个数组
	for ; i <= mid && j <= right; k++ {
		if nums[i] <= nums[j] {
			temp[k] = nums[i]
			i++
		} else {
			temp[k] = nums[j]
			j++
		}
	}

	// 虽然左右两边数组都已经插入到临时数组，但左右两边尾部的元素依旧没有插入，因此，这里还需要再插入一次
	// 左边继续合并
	for ; i <= mid; i++ {
		temp[k] = nums[i]
		k++
	}
	// 右边继续合并
	for ; j <= right; j++ {
		temp[k] = nums[j]
		k++
	}
	copy(nums[left:right+1], temp)
}

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	i := partition(nums, left, right)
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}
