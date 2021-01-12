package main

// 这道题类似于判断链表是否有环，可以使用哈希表和快慢指针
func isHappy(n int) bool {
	m := make(map[int]struct{})
	for n != 1 {
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = struct{}{}
		n = step(n)
	}
	return true
}

func isHappy1(n int) bool {
	slow, fast := n, step(n)
	for slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	return sum
}
