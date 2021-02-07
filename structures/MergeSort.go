package structures

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

func merge(nums []int, left, mid, right int) {
	tmp := make([]int, right-left+1) // 开辟一个临时数组，存放合并之后的数组
	i := left                        // 左边的指针，指向左边数组的第一个元素
	j := mid + 1                     // 右边的指针，指向右边数组的第一个元素
	k := 0                           // 用于记录临时数组当前指针指向
	// 一直遍历左右两个数组
	for ; i <= mid && j <= right; k++ {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
	}

	// 虽然左右两边数组都已经插入到临时数组，但左右两边尾部的元素依旧没有插入，因此，这里还需要再插入一次
	// 左边继续合并
	for ; i <= mid; i++ {
		tmp[k] = nums[i]
		k++
	}
	// 右边继续合并
	for ; j <= right; j++ {
		tmp[k] = nums[j]
		k++
	}
	copy(nums[left:right+1], tmp)
}
