package main

func isPalindrome(x int) bool {
	//首先所有的负数都不属于回文数
	//所有个位为 0 的数都不是回文数，因为反转后最高一位不可能是 0
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	num := 0
	for x > num {
		num = num*10 + x%10
		x /= 10
	}
	return x == num || num == x/10
}
