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
