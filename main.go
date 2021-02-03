package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func main() {
	fileNames := []string{}
	files, _ := ioutil.ReadDir("./leetcode-go")
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}
	fmt.Printf("目前为止总共做了 %v 道题\n", len(fileNames))
	n := getRandNumber(len(fileNames))
	fmt.Printf("生成的随机题目是 ----- %v\n", fileNames[n])
}

func getRandNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	i := patitation(nums, left, right)
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

func patitation(nums []int, left, right int) int {
	piovt := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < piovt {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}
