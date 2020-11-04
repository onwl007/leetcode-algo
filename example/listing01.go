package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)

	// 用来等待程序完成
	// 计数加 2，表示要等待两个 goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明一个匿名函数，并创建一个 goroutine
	go func() {
		defer wg.Done()

		for i := 0; i < 3; i++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		// 在函数退出时调用 Done 来通知 main 函数工作已经完成
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	// 等待 goroutine 结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
