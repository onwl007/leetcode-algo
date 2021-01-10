package main

// 维护一个哈希表，里面最多包含 k 个元素，当出现重复值时则说明在 k 距离内存在重复元素
// 每次遍历一个元素则将其加入哈希表中，如果哈希表的大小大于 k，则移除最前面的数字
// 类似滑动窗口，一直维护一个大小为 k 的滑动窗口一直往前走，判断窗口中是否有重复元素
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = i
		if len(m) > k {
			delete(m, nums[i-k])
		}
	}
	return false
}

// 先找到相等的元素，然后判断下标之差是否不超过 k
func containsNearbyDuplicate1(nums []int, k int) bool {
	m := make(map[int]int)

	for i, v := range nums {
		if j, ok := m[v]; ok && i-j <= k {
			return true
		}
		m[v] = i
	}
	return false
}
