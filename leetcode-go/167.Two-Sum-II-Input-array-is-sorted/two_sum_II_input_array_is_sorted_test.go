package main

import (
	"fmt"
	"testing"
)

type question167 struct {
	para167
	ans167
}

type para167 struct {
	one []int
	two int
}

type ans167 struct {
	one []int
}

func Test_Problem167(t *testing.T) {
	qs := []question167{

		{
			para167{[]int{2, 7, 11, 15}, 9},
			ans167{[]int{1, 2}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 167------------------------\n")

	for _, q := range qs {
		_, p := q.ans167, q.para167
		fmt.Printf("【input】:%v    【output】:%v\n", p.one, twoSum(p.one, p.two))
	}
	fmt.Printf("\n\n\n")
}
