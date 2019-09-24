package afGoNet

import (
	"bytes"
	"cxfProject/afGo/afGoface"
	"cxfProject/afGo/global"
	"encoding/binary"
	"errors"
)

type DataPack struct {
}

func NewDataPack() *DataPack {

	return &DataPack{}
}

//获取包的头的方法长度
func (dp *DataPack) GetHeadLen() uint32 {
	//dataLen uint32 4字节，Id uint32 4字节
	return 8
}

//封包方法
func (dp *DataPack) Pack(msg afGoface.IMessage) ([]byte, error) {

	dataBuff := bytes.NewBuffer([]byte{})

	//将dataLen 写进dataBuff中
	err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen())
	if err != nil {
		return nil, err
	}
	//将magId 写进dataBuff中
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId())
	if err != nil {
		return nil, err
	}

	//将data 写进dataBuff中
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetData())
	if err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

//拆包方法 ,将包的head信息读出来，之后再根据head信息里的data的长度在进行读一次
func (dp *DataPack) Unpack(binaryData []byte) (afGoface.IMessage, error) {

	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	//只解压head信息，得到dataLen和msgId

	msg := &Message{}

	err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen)
	if err != nil {
		return nil, err
	}

	err = binary.Read(dataBuff, binary.LittleEndian, &msg.Id)
	if err != nil {
		return nil, err
	}

	//判断dataLen 是否已经超出了我们允许的最大包长度
	if global.Cfg.MaxPackageSize > 0 && msg.DataLen > global.Cfg.MaxPackageSize {
		return nil, errors.New("too large dataSize!")
	}

	return msg, nil
}
