package main

import (
	"fmt"
	"testing"
)

type question26 struct {
	para26
	ans26
}

type para26 struct {
	one []int
}

type ans26 struct {
	one int
}

func Test_Problem26(t *testing.T) {
	qs := []question26{
		{
			para26{[]int{1, 1, 2}},
			ans26{2},
		},

		{
			para26{[]int{0, 0, 1, 1, 1, 1, 2, 3, 4, 4}},
			ans26{5},
		},

		{
			para26{[]int{0, 0, 0, 0, 0}},
			ans26{1},
		},

		{
			para26{[]int{1}},
			ans26{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 26------------------------\n")

	for _, q := range qs {
		_, p := q.ans26, q.para26
		fmt.Printf("[input]:%v      [output]:%v\n", p.one, removeDuplicates(p.one))
	}
	fmt.Printf("\n\n\n")
}
