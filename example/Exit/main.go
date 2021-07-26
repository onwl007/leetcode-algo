package main

import (
	"fmt"
	"os"
)

func main() {
	// 使用 os.Exit 时 defer 将不会被执行，所以这里的 fmt.Println 将永远不会被调用
	defer fmt.Println("!")

	os.Exit(3)
}
