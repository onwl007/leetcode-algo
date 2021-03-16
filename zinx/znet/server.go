package znet

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

func (s *Server) Start() {}

func (s *Server) Stop() {}

func (s *Server) Serve() {}
