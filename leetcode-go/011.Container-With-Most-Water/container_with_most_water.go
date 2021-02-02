package main

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	area, start, end := 0, 0, len(height)-1
	for start <= end {
		width := end - start
		high := 0
		if height[start] < height[end] {
			high = height[start]
			start++
		} else {
			high = height[end]
			end--
		}

		tmp := width * high
		if tmp > area {
			area = tmp
		}
	}
	return area
}

// @lc code=end

func maxArea1(height []int) int {
	area := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			area = max(area, min(height[i], height[j])*(j-i))
		}
	}
	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
