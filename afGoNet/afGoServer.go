package afGoNet

import (
	"cxfProject/afGo/afGoface"
	"cxfProject/afGo/global"
	"errors"
	"fmt"
	"net"
)

type AfGoServer struct {
	Name      string
	IpVersion string
	Ip        string
	Port      int

	//当前的server的消息管理模块，用来绑定msgId和对应的处理业务api关系
	MsgHandler afGoface.IMsgHandle

	//该server的连接管理器
	ConnManage afGoface.IConnectionManage

	//启动后调用的Hook函数

	OnConnStart func(conn afGoface.IConnection)

	//断开时候调用的hook函数
	OnConnStop func(conn afGoface.IConnection)
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

func (s *AfGoServer) AddRouter(msgId uint32, router afGoface.IRouter) {

	s.MsgHandler.AddRouter(msgId, router)
	fmt.Println("add router success!")
}
func (s *AfGoServer) Start() {

	fmt.Println("start... [afGo]")

	go func() {
		s.MsgHandler.StartWorkerPool()
		//获取tcp的addr
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

			//设置最大连接数数量判断，如果超过最大连接数量，那么关闭新建的连接

			if s.ConnManage.GetConnectionLen() >= global.Cfg.MaxConn {
				//todo 给客户端回应一个错误
				conn.Close()
				continue
			}

			//处理新链接的业务方法和conn进行绑定 得到我们的链接模块
			dealConn := NewConnection(s, conn, cid, s.MsgHandler)

			cid++
			go dealConn.Start()
		}

	}()

	//监听服务器的地址

	//阻塞的等待客户端进行连接，处理客户端链接业务
}

func (s *AfGoServer) GetConnMange() afGoface.IConnectionManage {

	return s.ConnManage
}

func (s *AfGoServer) Stop() {

	fmt.Println("[server stop] ", s.Name)
	s.ConnManage.ClearConnection()

}
func (s *AfGoServer) Server() {
	//启动server的功能
	s.Start()

	//todo 做一些启动之外的功能
	select {}

}

func NewServer(name string) afGoface.IServerFace {

	s := &AfGoServer{
		Name:       global.Cfg.Name,
		IpVersion:  "tcp4",
		Ip:         global.Cfg.Host,
		Port:       global.Cfg.TcpPort,
		MsgHandler: NewMessageHandle(),
		ConnManage: NewConnManage(),
	}

	return s
}

//注册hookStart方法

func (s *AfGoServer) SetOnConnStart(hookFun func(conn afGoface.IConnection)) {
	s.OnConnStart = hookFun
}

//注册hookStop方法
func (s *AfGoServer) SetOnConnStop(hookFun func(conn afGoface.IConnection)) {

	s.OnConnStop = hookFun
}

//调用hookStart方法
func (s *AfGoServer) CallOnConnStart(conn afGoface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("Call OnConnStart")
		s.OnConnStart(conn)
	}
}

//调用hookStop方法
func (s *AfGoServer) CallOnConnStop(conn afGoface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("Call OnConnStop")
		s.OnConnStart(conn)
	}
}
