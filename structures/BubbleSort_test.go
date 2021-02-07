package structures

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Case struct {
	nums []int
	out  []int
}

func TestBubbleSort(t *testing.T) {
	qs := []Case{
		{nums: []int{1, 3, 2, 6, 9, 4, 8}, out: []int{1, 3, 2, 6, 9, 4, 8}},
		{nums: []int{1}, out: []int{1}},
		{nums: []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, out: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	ast := assert.New(t)
	for _, v := range qs {
		sort.Ints(v.out)
		BubbleSort(v.nums)
		ast.Equal(v.out, v.nums)
	}
}
