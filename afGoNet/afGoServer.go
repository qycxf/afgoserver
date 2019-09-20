package afGoNet

import (
	"afGo/afGoface"
	"afGo/global"
	"errors"
	"fmt"
	"net"
)

type AfGoServer struct {
	Name      string
	IpVersion string
	Ip        string
	Port      int

	//当前的server添加一个router，server注册的链接对应的处理业务

	Router afGoface.IRouter
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {

	//回显业务

	fmt.Println("[conn Handle] CallBackToClient")

	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)

		return errors.New("CallBackToClient err")
	}

	return nil

}

func (s *AfGoServer) AddRouter(router afGoface.IRouter) {

	s.Router = router
	fmt.Println("add router success!")
}
func (s *AfGoServer) Start() {

	fmt.Println("start... [afGo]")
	//获取tcp的addr

	go func() {

		fmt.Println(s.Ip)
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
		var cid uint32
		cid = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("accept err", err)
				continue
			}

			//已经与客户端建立链接，做一些业务
			//处理新链接的业务方法和conn进行绑定 得到我们的链接模块
			dealConn := NewConnection(conn, cid, s.Router)

			cid++
			go dealConn.Start()
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
		Name:      global.Cfg.Name,
		IpVersion: "tcp4",
		Ip:        global.Cfg.Host,
		Port:      global.Cfg.TcpPort,
		Router:    nil,
	}

	return s
}
