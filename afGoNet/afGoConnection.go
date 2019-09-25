package afGoNet

import (
	"cxfProject/afGo/afGoface"
	"errors"
	"fmt"
	"io"
	"net"
)

type Connection struct {
	//当前链接的socket TCP套接字
	Conn *net.TCPConn

	//链接的id
	ConnID uint32

	//当前链接的状态
	IsClose bool

	//告知当前链接已经退出的/停止 channel
	ExitChan chan bool

	//消息的管理msgId和对应的处理业务api关系

	MsgHandler afGoface.IMsgHandle
}

func NewConnection(conn *net.TCPConn, connID uint32,
	handle afGoface.IMsgHandle) *Connection {

	c := &Connection{

		Conn:       conn,
		ConnID:     connID,
		MsgHandler: handle,
		IsClose:    false,
		ExitChan:   make(chan bool, 1),
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
		//buf := make([]byte, global.Cfg.MaxPackageSize)
		//_, err := c.Conn.Read(buf)
		//
		//if err != nil {
		//	fmt.Println("recv buf err", err)
		//
		//	continue
		//}

		//创建拆包、解包对象
		dp := NewDataPack()

		//读取客户端的msd head 二进制流 8个字节
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), headData)
		if err != nil {
			fmt.Println("read dataHead err", err)
			break
		}
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack err", err)
			break
		}
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			_, err := io.ReadFull(c.GetTCPConnection(), data)

			if err != nil {
				fmt.Println("read msg data err", err)
			}
		}
		msg.SetData(data)

		//拆包，得到msgId 和dataLen 放在msg消息中

		//根据dataLen 再次读取data 放在msg data中
		//得到当前conn的数据的request请求数据

		req := Request{
			conn: c,
			msg:  msg,
		}

		go c.MsgHandler.DoMsgHandler(&req)

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

//send message 方法，将要发送给客户端的数据先进行封装
func (c *Connection) SendMsg(msgId uint32, data []byte) error {

	if c.IsClose {
		return errors.New("connection is close")
	}

	//将data封包

	dp := NewDataPack()

	binaryMsg, err := dp.Pack(&Message{
		Id:   msgId,
		Data: data,
	})

	if err != nil {
		fmt.Println("Pack error msg")
		return errors.New("Pack error")
	}

	_, err = c.Conn.Write(binaryMsg)

	if err != nil {
		fmt.Println("write msg id", msgId, "error:", err)

		return errors.New("conn write error")
	}
	return nil
}
