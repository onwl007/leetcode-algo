package znet

import (
	"fmt"
	"net"

	"github.com/onwl007/leetcode-algo/zinx/ziface"
)

// iServer的接口实现，定义一个 Server 的服务器模块
type Server struct {
	// 服务器名称
	Name string
	// 服务器绑定的 IP 版本
	IPVersion string
	// 服务器监听的 IP
	IP string
	// 服务器监听的端口
	Port int
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listener at IP: %s, Port %d, is starting\n", s.IP, s.Port)

	go func() {
		// 1 获取 TCP 的 Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		// 2 监听服务器的地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, " err ", err)
			return
		}

		fmt.Println("start Zinx server success ", s.Name, "success, Listenning...")

		// 3 阻塞的等待客户端连接，处理客户端连接业务（读写）
		for {
			// 如果有客户端连接过来，阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			// 已经与客户端建立连接，做一些业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err", err)
						continue
					}

					fmt.Printf("recv client buf %s, cnt %d\n", buf, cnt)
					// 回显功能
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {}

func (s *Server) Serve() {
	// 启动 server 的服务功能
	s.Start()

	// 阻塞状态
	select {}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
