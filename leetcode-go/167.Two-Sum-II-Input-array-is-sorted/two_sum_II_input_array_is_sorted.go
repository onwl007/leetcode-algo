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
