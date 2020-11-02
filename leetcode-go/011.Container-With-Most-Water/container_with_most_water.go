package main

func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			width := j - i
			tmp := width * min(height[i], height[j])
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
