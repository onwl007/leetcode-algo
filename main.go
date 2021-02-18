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

	getTencentProblem()
}

func getTencentProblem() {
	n := getRandNumber(len(TencentProblems))
	fmt.Printf("生成的随机题目是 ----- %v.%v\n", TencentProblems[n].Num, TencentProblems[n].Title)
}

func getRandNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

type TencentProblem struct {
	Num   int
	Title string
}

var TencentProblems = []TencentProblem{
	{Num: 2, Title: "两数相加"},
	{Num: 4, Title: "寻找两个正序数组的中位数"},
	{Num: 5, Title: "最长回文子串"},
	{Num: 7, Title: "整数反转"},
	{Num: 8, Title: "字符串转换整数 (atoi)"},
	{Num: 9, Title: "回文数"},
	{Num: 11, Title: "盛最多水的容器"},
	{Num: 14, Title: "最长公共前缀"},
	{Num: 15, Title: "三数之和"},
	{Num: 16, Title: "最接近的三数之和"},
	{Num: 20, Title: "有效的括号"},
	{Num: 21, Title: "合并两个有序链表"},
	{Num: 23, Title: "合并K个升序链表"},
	{Num: 26, Title: "删除排序数组中的重复项"},
	{Num: 33, Title: "搜索旋转排序数组"},
	{Num: 43, Title: "字符串相乘"},
	{Num: 46, Title: "全排列"},
	{Num: 53, Title: "最大子序和"},
	{Num: 54, Title: "螺旋矩阵"},
	{Num: 59, Title: "螺旋矩阵 II"},
	{Num: 61, Title: "旋转链表"},
	{Num: 62, Title: "不同路径"},
	{Num: 70, Title: "爬楼梯"},
	{Num: 78, Title: "子集"},
	{Num: 88, Title: "合并两个有序数组"},
	{Num: 89, Title: "格雷编码"},
	{Num: 104, Title: "二叉树的最大深度"},
	{Num: 121, Title: "买卖股票的最佳时机"},
	{Num: 122, Title: "买卖股票的最佳时机 II"},
	{Num: 124, Title: "二叉树中的最大路径和"},
	{Num: 136, Title: "只出现一次的数字"},
	{Num: 141, Title: "环形链表"},
	{Num: 142, Title: "环形链表 II"},
	{Num: 146, Title: "LRU 缓存机制"},
	{Num: 148, Title: "排序链表"},
	{Num: 155, Title: "最小栈"},
	{Num: 160, Title: "相交链表"},
	{Num: 169, Title: "多数元素"},
	{Num: 206, Title: "反转链表"},
	{Num: 215, Title: "数组中的第K个最大元素"},
	{Num: 217, Title: "存在重复元素"},
	{Num: 230, Title: "二叉搜索树中第K小的元素"},
	{Num: 231, Title: "2的幂"},
	{Num: 235, Title: "二叉搜索树的最近公共祖先"},
	{Num: 236, Title: "二叉树的最近公共祖先"},
	{Num: 237, Title: "删除链表中的节点"},
	{Num: 238, Title: "除自身以外数组的乘积"},
	{Num: 292, Title: "Nim 游戏"},
	{Num: 344, Title: "反转字符串"},
	{Num: 557, Title: "反转字符串中的单词 III"},
}
