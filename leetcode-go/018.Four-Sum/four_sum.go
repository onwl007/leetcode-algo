package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums); i++ {
		sums := target - nums[i]
		if i == 0 || nums[i] != nums[i-1] {
			for j := 1; j < len(nums); j++ {
				l, r := j+1, len(nums)-1
				if j == 0 || nums[j] != nums[j-1] {
					for l < r {
						if nums[j]+nums[l]+nums[r] == sums {
							res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
							for l < r && nums[l] == nums[l+1] {
								l++
							}
							for l < r && nums[r] == nums[r-1] {
								r--
							}
							l++
							r--
						} else if nums[j]+nums[l]+nums[r] < sums {
							l++
						} else {
							r--
						}
					}
				}
			}
		}
	}
	return res
}
