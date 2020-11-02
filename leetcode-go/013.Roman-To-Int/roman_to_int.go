package main

func romanToInt(s string) int {
	pre, r := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := getValue(s[i])
		if cur >= pre {
			r += cur
		} else {
			r -= cur
		}
		pre = cur
	}
	return r
}

func getValue(b byte) int {
	v := 0
	switch b {
	case 'I':
		v = 1
	case 'V':
		v = 5
	case 'X':
		v = 10
	case 'L':
		v = 50
	case 'C':
		v = 100
	case 'D':
		v = 500
	case 'M':
		v = 1000
	default:
		v = 0
	}
	return v
}
