package main

func mySqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := left + (right-left)/2
		target := mid * mid
		if target == x {
			return mid
		} else if target > x {
			right = mid - 1
		} else if target < x {
			left = mid + 1
		}
	}
	return right
}
