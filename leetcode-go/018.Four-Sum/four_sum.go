package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	n := len(nums)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 计算当前的最小值，如果最小值都会比 target 大，后面就不用再继续计算了
		min := nums[i] + nums[i+1] + nums[i+2] + nums[i+3]
		if min > target {
			continue
		}
		// 计算当前的最大值，如果最大值都比 target 小，后面就不用再继续计算了
		max := nums[i] + nums[n-1] + nums[n-2] + nums[n-3]
		if max < target {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			// 相邻元素重复，会产生重复结果
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			min1 := nums[i] + nums[j] + nums[j+1] + nums[j+2]
			if min1 > target {
				break
			}
			max1 := nums[i] + nums[j] + nums[n-1] + nums[n-2]
			if max1 < target {
				break
			}
			l, r := j+1, n-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}
