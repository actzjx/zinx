package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Server struct {
	// 服务器的名称
	Name string
	// 服务器绑定的ip版本
	IPVersion string
	// 服务器监听的IP
	IP string
	// 服务器监听的端口
	Port uint
}

func (s *Server) Start() {
	fmt.Printf("[start] %s Listener at IP: %s, Port: %d, is starting \n", s.Name, s.IP, s.Port)
	// 1 获取TCP的Addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error", err)
		return
	}
	// 2 监听服务器的地址
	listenner, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("listen", s.IPVersion, " error ", err)
	}
	fmt.Println("start Zinx server success ", s.Name)
	// 3 阻塞的等待客户端链接，处理客户端链接业务（读写）
	for {
		// 如果有客户端链接过来，阻塞会返回
		conn, err := listenner.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err", err)
			continue
		}

		// 客户端建立链接，做一些业务，做一个最大512字节长度的回西显业务
		go func() {
			for {
				buf := make([]byte, 512)
				cnt, err := conn.Read(buf)
				if err != nil {
					fmt.Println("recv buf err ", err)
					continue
				}

				fmt.Printf("receive client buf %s, cnt %d\n", buf, cnt)
				// 回显
				if _, err := conn.Write(buf[:cnt]); err != nil {
					fmt.Println("write back buffer err ", err)
					continue
				}

			}
		}()
	}
}

func (s *Server) Stop() {
	//TODO 将一些服务器资源、状态停止和回收
}

func (s *Server) Serve() {
	go s.Start()

	// 阻塞
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
