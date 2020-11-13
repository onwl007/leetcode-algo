package main

import (
	"fmt"
	"testing"
)

type question704 struct {
	para704
	ans704
}

type para704 struct {
	nums   []int
	target int
}

type ans704 struct {
	one int
}

func Test_Problem704(t *testing.T) {
	qs := []question704{

		{
			para704{[]int{-1, 0, 3, 5, 9, 12}, 9},
			ans704{4},
		},

		{
			para704{[]int{-1, 0, 3, 5, 9, 12}, 2},
			ans704{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 704------------------------\n")

	for _, q := range qs {
		_, p := q.ans704, q.para704
		fmt.Printf("[input]:%v      [output]:%v\n", p, search(p.nums, p.target))
	}
	fmt.Printf("\n\n\n")
}
