package main

import "math"

func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int)
	for i, v := range list1 {
		m[v] = i
	}

	var vals []string
	minSum := math.MaxInt32
	for i, v := range list2 {
		if _, ok := m[v]; ok {
			sum := m[v] + i
			if sum < minSum {
				vals = []string{}
				minSum = sum
				vals = append(vals, v)
			} else if sum == minSum {
				vals = append(vals, v)
			}
		}
	}
	return vals
}
