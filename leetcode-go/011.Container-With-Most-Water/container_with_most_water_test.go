package main

import (
	"fmt"
	"testing"
)

type question11 struct {
	para11
	ans11
}

type para11 struct {
	one []int
}

type ans11 struct {
	one int
}

func Test_Problem11(t *testing.T) {
	qs := []question11{
		{
			para11{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			ans11{49},
		},
		{
			para11{[]int{1, 1}},
			ans11{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 11------------------------\n")

	for _, q := range qs {
		_, p := q.ans11, q.para11
		fmt.Printf("[input]: %v      [output]: %v\n", p.one, maxArea(p.one))
	}
	fmt.Printf("\n\n\n")
}
