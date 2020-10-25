package main

import "strconv"

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	length := len(str)
	for i := range str {
		if str[length-i-1] != str[i] {
			return false
		}
	}
	return true
}
