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

func bubbleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		flag := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				tmp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = tmp
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

func insertSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > val {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = val
	}
}

func MergeSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	mergeSort(nums, left, mid)
	mergeSort(nums, mid+1, right)
	merge(nums, left, mid, right)
}

func merge(nums []int, left, mid, right int) {
	tmp := []int{}
	i := left
	j := mid + 1
	k := 0
	for ; i <= mid && j <= right; k++ {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		tmp[k] = nums[i]
		k++
	}
	for ; j <= right; j++ {
		tmp[k] = nums[j]
		k++
	}
	copy(nums[left:right+1], tmp)
}

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	i := partition(nums, left, right)
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}
