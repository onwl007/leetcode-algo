package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

// go tool trace trace.out
func traceTool() {
	// 1. 创建一个 trace 文件
	f, err := os.Create("./trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// 2. 启动 trace
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	// 正常调试的业务
	fmt.Println("Hello GMP")

	// 3. 停止 trace
	trace.Stop()
}
