package afGoNet

import (
	"afGo/afGoface"
	"fmt"
	"net"
)

type Connection struct {
	//当前链接的socket TCP套接字
	Conn *net.TCPConn

	//链接的id
	ConnID uint32

	//当前链接的状态
	IsClose bool

	//当前链接所绑定的处理业务的方法API
	handleAPI afGoface.HandleFunc

	//告知当前链接已经退出的/停止 channel
	ExitChan chan bool

	//该链接处理的方法

	Router afGoface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32,
	callBack_api afGoface.HandleFunc) *Connection {

	c := &Connection{

		Conn:      conn,
		ConnID:    connID,
		handleAPI: callBack_api,
		IsClose:   false,
		ExitChan:  make(chan bool, 1),
	}

	return c
}

func (c *Connection) Stop() {

	fmt.Println("conn stop...ConnId=", c.ConnID)

	if c.IsClose {
		return
	}

	//关闭socket链接
	c.IsClose = true
	c.Conn.Close()

	close(c.ExitChan)

}

//链接的读业务方法
func (c *Connection) StartReader() {

	fmt.Println("reader Goroutine is running...")

	defer fmt.Println("connId=", c.ConnID, "reader is exit")

	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)

		if err != nil {
			fmt.Println("recv buf err", err)

			continue
		}
		//得到当前conn的数据的request请求数据

		req := Request{
			conn: c,
			data: buf,
		}

		go func(request afGoface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

		//从路由中，找到注册的绑定的conn对应的router调用

	}

}
func (c *Connection) Start() {

	fmt.Println("conn start... connId=", c.ConnID)

	//启动从当前链接的读取数据业务
	go c.StartReader()
	//todo 启动从当前链接写数据的业务

}
func (c *Connection) GetTCPConnection() *net.TCPConn {

	return c.Conn
}
func (c *Connection) GetConnID() uint32 {

	return c.ConnID
}

func (c *Connection) GetRemoteAddr() net.Addr {

	return c.Conn.RemoteAddr()

}
func (c *Connection) Send(data []byte) error {
	return nil
}
