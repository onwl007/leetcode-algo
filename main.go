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

	arr := []int{4, 5, 6, 3, 2, 1}
	bubbleSort(arr)
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
		// 提前退出冒泡循环的标志位
		flag := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
				flag = true // 表示有数据交换
			}
		}
		if !flag {
			break // 没有数据交换，提前退出
		}
	}
}
