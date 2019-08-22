package afGoNet

import (
	"afGo/afGoface"
	"fmt"
	"net"
)

type AfGoServer struct {
	Name      string
	IpVersion string
	Ip        string

	Port int
}

func (s *AfGoServer) Start() {

	fmt.Println("start... afGo")
	//获取tcp的addr

	go func() {
		addr, err := net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d", s.Ip, s.Port))

		if err != nil {
			fmt.Println("start error:", err)
			return
		}

		listener, err := net.ListenTCP(s.IpVersion, addr)
		if err != nil {
			fmt.Println("listener err:", err)
			return
		}

		fmt.Println("start afGo server success,", s.Name)

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept err", err)
				continue
			}

			//已经与客户端建立链接，做一些业务

			go func() {
				for {
					buf := make([]byte, 512)

					cnt, err := conn.Read(buf)

					if err != nil {
						fmt.Println("receive buf err", err)
						continue
					}

					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}
				}

			}()
		}

	}()

	//监听服务器的地址

	//阻塞的等待客户端进行连接，处理客户端链接业务
}

func (s *AfGoServer) Stop() {

}
func (s *AfGoServer) Server() {
	//启动server的功能
	s.Start()

	//todo 做一些启动之外的功能
	select {}

}

func NewServer(name string) afGoface.IServerFace {

	s := &AfGoServer{
		Name:      name,
		IpVersion: "tcp4",
		Ip:        "0.0.0.0",
		Port:      8999,
	}

	return s
}
