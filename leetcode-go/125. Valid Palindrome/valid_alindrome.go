package main

import "strings"

func isPalindrome(s string) bool {
	str := ""
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			str += string(s[i])
		}
	}

	l, r := 0, len(str)-1
	str = strings.ToLower(str)
	for l < r {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isalnum(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}
