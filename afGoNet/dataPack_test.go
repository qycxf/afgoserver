package afGoNet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

//测试dataPack拆包、封包 单元测试
func TestDtaPack(t *testing.T) {

	//模拟的服务器

	//1.创建socketTcp

	listenner, err := net.Listen("tcp", "0.0.0.0:7777")

	if err != nil {
		return
	}

	//创建一个go 负责送客户端处理业务
	go func() {
		for {
			conn, err := listenner.Accept()

			if err != nil {
				fmt.Println("server accept err", err)
				return
			}

			go func(conn net.Conn) {
				//处理客户端的请求
				dp := NewDataPack()
				//拆包的过程
				for {
					//1.读取head
					headData := make([]byte, dp.GetHeadLen())
					//读满
					_, err := io.ReadFull(conn, headData)

					if err != nil {
						break
					}

					msgHead, err := dp.Unpack(headData)
					if err != nil {
						return
					}
					//msg 有数据，需要进行第二次读取
					if msgHead.GetMsgLen() > 0 {
						//2.根据head中的dataLen 读取内容
						msg := msgHead.(*Message)

						msg.Data = make([]byte, msg.GetMsgLen())

						//根据dataLen的长度在此从io流中读取

						_, err := io.ReadFull(conn, msg.Data)

						if err != nil {
							return
						}

						fmt.Println("msg:", msg)
					}
				}
			}(conn)
		}

	}()

	//模拟客户端
	conn, err := net.Dial("tcp", "0.0.0.0:7777")

	if err != nil {
		return
	}

	//创建一个封包对象

	dp := NewDataPack()

	//模拟粘包过程,封装两个msg一同发送

	//封装第一个
	msg1 := &Message{
		Id:      1,
		Data:    []byte{'z', 'i', 'n', 'x'},
		DataLen: 4,
	}

	sendData1, err := dp.Pack(msg1)

	if err != nil {
		return
	}
	//封装第二个
	msg2 := &Message{
		Id:      2,
		Data:    []byte{'n', 'i', 'h', 'a', 'o', '!', '!'},
		DataLen: 7,
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		return
	}
	sendData1 = append(sendData1, sendData2...)
	conn.Write(sendData1)
	select {}

}
