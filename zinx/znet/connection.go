package znet

import (
	"net"

	"github.com/onwl007/leetcode-algo/zinx/ziface"
)

// 连接模块
type Connection struct {
	// 当前连接的 socket TCP 套接字
	Conn *net.TCPConn

	// 连接的 ID
	ConnID uint32

	// 当前的连接状态
	isClosed bool

	// 当前连接所绑定的处理业务方法 API
	handleAPI ziface.HandleFunc

	// 告知当前连接已经退出的、停止 channel
	ExitChan chan bool
}

// 初始化连接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandleFunc) *Connection  {
	c := &Connection{
		Conn: conn,
		ConnID: connID,
		handleAPI: callback_api,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}

	return c
}