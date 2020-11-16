package main

// 使用二分查找，先固定一个数，在这个数的右边使用二分查找找出符合的元素
func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		left, right := i+1, len(numbers)-1
		for left <= right {
			mid := left + (right-left)/2
			t := target - numbers[i]
			if t == numbers[mid] {
				return []int{i + 1, mid + 1}
			} else if t > numbers[mid] {
				left = mid + 1
			} else if t < numbers[mid] {
				right = mid - 1
			}
		}
	}
	return []int{-1, -1}
}

// 双指针法，定义左右两个指针，判断两个的和是否等于 target
// 数组有序
// 当和大于 target 时，说明右指针指向的值比较大，需要左移
// 当和小于 target 时，说明左指针指向的值比较小，需要右移
func twoSum1(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left <= right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else if sum < target {
			left++
		}
	}
	return []int{-1, -1}
}
