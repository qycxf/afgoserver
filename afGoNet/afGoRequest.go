package afGoNet

import "cxfProject/afGo/afGoface"

type Request struct {
	conn afGoface.IConnection

	msg afGoface.IMessage
}

func (r *Request) GetConnection() afGoface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {

	return r.msg.GetData()
}
func (r *Request) GetMsgId() uint32 {

	return r.msg.GetMsgId()
}
