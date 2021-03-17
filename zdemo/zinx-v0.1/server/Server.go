package main

import "github.com/onwl007/leetcode-algo/zinx/znet"

func main() {
	// 1 创建一个 server 句柄，使用 Zinx 的 api
	s := znet.NewServer("[zinx v0.1]")
	// 2 启动 server
	s.Serve()
}
